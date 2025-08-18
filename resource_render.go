package apui

import (
	"bytes"
	"encoding/json"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"reflect"
	"slices"
)

type ResourceType int

const (
	Unknown ResourceType = iota
	Entity
	Collection
)

type ResourceTypeDetector interface {
	DetectResourceType(response any) ResourceType
	CollectionItems(response any) any
}

type resourceTypeDetector struct{}

func (r *resourceTypeDetector) DetectResourceType(response any) ResourceType {
	t := reflect.TypeOf(response)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Struct, reflect.Map:
		return Entity
	case reflect.Slice, reflect.Array:
		return Collection
	}
	return Unknown
}

func (r *resourceTypeDetector) CollectionItems(response any) any {
	if r.DetectResourceType(response) == Collection {
		return response
	}
	return nil
}

var defaultResourceTypeDetector ResourceTypeDetector = &resourceTypeDetector{}

func (b *Browser) writeMain(ctx aitch.ImperativeContext) error {
	if res, ok := getContextResponse(ctx.Context()); ok {
		t := b.resourceTypeDetector.DetectResourceType(res)
		switch t {
		case Entity:
			b.writeMainEntity(ctx, res)
		case Collection:
			b.writeMainCollection(ctx, res)
		default:
			b.writeMainUnknown(ctx, res)
		}
	}
	return nil
}

var (
	classKeep       = html.Class("keep")
	classRight      = html.Class("right")
	classNull       = html.Class("null")
	tableRenderNode = html.Table(
		html.THead(
			html.Tr(
				aitch.IterateKey("headers", html.Th(func(ctx *context.Context) []byte {
					return []byte(ctx.Cargo.(string))
				})),
				html.Th(
					html.Em("(raw JSON)"),
				),
			),
		),
		html.TBody(
			aitch.IterateKey("items", rowRenderNode),
		),
	)
	rowRenderNode = html.Tr(
		aitch.Imperative(func(ctx aitch.ImperativeContext) error {
			row := ctx.Context().Data
			if hdrs, ok := context.Get[[]string](ctx.Context().Parent, "headers"); ok {
				for _, hdr := range hdrs {
					v, present := row[hdr]
					renderColumn(ctx, present, hdr, v, false)
				}
			}
			ctx.Start(elemTd, false).
				Start(elemDetails, false, classKeep).
				Start(elemSummary, false).WriteString("JSON").End()
			_ = jsonRenderNode.Render(context.New(ctx.Context().Writer, nil, row))
			ctx.EndAll()
			return nil
		}),
	)
)

func renderColumn(ctx aitch.ImperativeContext, present bool, name string, value any, noAlign bool) {
	if !present {
		ctx.Start(elemTd, false).End()
		return
	} else if value == nil {
		ctx.Start(elemTd, false).
			Start(elemSpan, false, classNull).
			WriteString("null").
			End().End()
		return
	}
	var att aitch.Node
	switch tv := value.(type) {
	case json.Number:
		att = classRight
		value = tv.String()
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		att = classRight
	case bool:
		// nothing special
	case string:
		if DefaultUriPropertyDetector != nil && DefaultUriPropertyDetector.IsUriProperty(name) {
			ctx.Start(elemTd, false).
				Start(elemA, false, html.Href(tv)).
				WriteString(tv).
				End().End()
			return
		}
	default:
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice, reflect.Array:
			renderColumnJson(ctx, value, "Array")
			return
		case reflect.Map, reflect.Struct:
			renderColumnJson(ctx, value, "Object")
			return
		}
	}
	if noAlign {
		att = nil
	}
	ctx.Start(elemTd, false, att).
		WriteNodes(aitch.Text(value)).End()
}

func renderColumnJson(ctx aitch.ImperativeContext, value any, title string) {
	ctx.Start(elemTd, false).
		Start(elemDetails, false, classKeep).
		Start(elemSummary, false).WriteString(title).End()
	_ = jsonRenderNode.Render(context.New(ctx.Context().Writer, nil, value))
	ctx.End().End()
}

func (b *Browser) writeMainEntity(ctx aitch.ImperativeContext, response any) {
	open := html.Open()
	var obj map[string]any
	isObj := false
	switch tv := response.(type) {
	case map[string]any:
		obj = tv
		isObj = true
	default:
		t := reflect.TypeOf(response)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() == reflect.Struct {
			if data, err := json.Marshal(response); err == nil {
				if err = json.Unmarshal(data, &obj); err == nil {
					isObj = true
				}
			}
		}
	}
	_ = obj
	if isObj {
		open = nil
	}
	ctx.Start(elemDiv, false).
		Start(elemDetails, false, classKeep, open).
		Start(elemSummary, false).WriteString("JSON").End()
	_ = jsonRenderNode.Render(context.New(ctx.Context().Writer, nil, response))
	ctx.End().End()
	if isObj {
		ptys := make([]string, 0)
		for k := range obj {
			ptys = append(ptys, k)
		}
		slices.Sort(ptys)
		ctx.Start(elemDiv, false).
			Start(elemDetails, false, classKeep, html.Open()).
			Start(elemSummary, false).WriteString("Tabular").End().
			Start(elemTable, false, html.Class("tabular")).
			Start(elemTHead, false).
			Start(elemTr, false).
			Start(elemTh, false).WriteString("property").End().
			Start(elemTh, false, html.Class("value")).WriteString("value").End().
			End(). //tr
			End(). //thead
			Start(elemTBody, false)
		for _, pty := range ptys {
			ctx.Start(elemTr, false).
				Start(elemTh, false, classRight).
				WriteString(pty).End()
			renderColumn(ctx, true, pty, obj[pty], true)
			ctx.End() //tr
		}
		ctx.End().End().End().End()
	}
}

func (b *Browser) writeMainCollection(ctx aitch.ImperativeContext, response any) {
	c := b.convertToCollection(b.resourceTypeDetector.CollectionItems(response))
	hdrsSeen := make(map[string]struct{})
	hdrs := make([]string, 0)
	for _, item := range c {
		for k := range item {
			if _, ok := hdrsSeen[k]; !ok {
				hdrs = append(hdrs, k)
			}
			hdrsSeen[k] = struct{}{}
		}
	}
	slices.Sort(hdrs)
	tableCtx := &context.Context{
		Data: map[string]any{
			"headers": hdrs,
			"items":   c,
		},
		Writer: ctx.Context().Writer,
		Parent: ctx.Context(),
	}
	_ = tableCtx
	_ = tableRenderNode.Render(context.New(ctx.Context().Writer, map[string]any{
		"headers": hdrs,
		"items":   c,
	}, nil))
}

func (b *Browser) convertToCollection(response any) []map[string]any {
	if tr, ok := response.([]map[string]any); ok {
		return tr
	}
	result := make([]map[string]any, 0)
	if data, err := json.Marshal(response); err == nil {
		d := json.NewDecoder(bytes.NewReader(data))
		d.UseNumber()
		_ = d.Decode(&result)
	}
	return result
}

func (b *Browser) writeMainUnknown(ctx aitch.ImperativeContext, response any) {
	jctx := &context.Context{
		Cargo:  ctx.Context().Data["response"],
		Writer: ctx.Context().Writer,
		Parent: ctx.Context(),
	}
	_ = b.jsonRenderer.Render(jctx)
}
