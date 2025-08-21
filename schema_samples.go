package apui

import (
	"github.com/go-andiamo/chioas"
	"reflect"
	"strings"
)

func (b *Browser) methodRequestSample(m chioas.Method) (sample string, has bool) {
	if has = m.Request != nil; has {
		if schema := b.methodRequestSchema(m.Request); schema != nil {
			sample = schemaSample(schema, m.Request.IsArray)
		} else {
			sample = "{\n}"
		}
	}
	return sample, has
}

func (b *Browser) methodRequestSchema(mr *chioas.Request) *chioas.Schema {
	req := mr
	if mr.Ref != "" {
		if b.definition.Components != nil {
			// lookup in components...
			if r, ok := b.definition.Components.Requests[mr.Ref]; ok {
				req = &r
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	if req.SchemaRef != "" {
		if b.definition.Components == nil {
			return nil
		}
		for _, s := range b.definition.Components.Schemas {
			if s.Name == req.SchemaRef {
				return &s
			}
		}
	} else if req.Schema != nil {
		return schemaFrom(req.Schema)
	}
	return nil
}

func schemaFrom(s any) *chioas.Schema {
	switch st := s.(type) {
	case *chioas.Schema:
		return st
	case chioas.Schema:
		return &st
	default:
		t := reflect.TypeOf(s)
		switch t.Kind() {
		case reflect.Struct:
			sf := &chioas.Schema{}
			if schema, err := sf.From(s); err == nil {
				return schema
			}
		}
	}
	return nil
}

func schemaSample(schema *chioas.Schema, isArray bool) string {
	var builder strings.Builder
	indent := strings.Repeat(" ", JsonIndent)
	if isArray {
		builder.WriteString("[\n    {\n")
		indent = strings.Repeat(" ", JsonIndent*2)
	} else {
		builder.WriteString("{\n")
	}
	schemaSampleProperties(&builder, indent, schema.Properties)
	if isArray {
		builder.WriteString("    }\n]")
	} else {
		builder.WriteString("}")
	}
	return builder.String()
}

func schemaSampleProperties(builder *strings.Builder, indent string, ptys chioas.Properties) {
	for i, pty := range ptys {
		builder.WriteString(indent)
		builder.WriteString(`"` + pty.Name + `": `)
		switch pty.Type {
		case "integer":
			builder.WriteString("0")
		case "number":
			builder.WriteString("0.0")
		case "boolean":
			builder.WriteString("false")
		case "array":
			builder.WriteString("[]")
		case "object":
			if len(pty.Properties) > 0 {
				builder.WriteString("{\n")
				schemaSampleProperties(builder, indent+strings.Repeat(" ", JsonIndent), pty.Properties)
				builder.WriteString(indent)
				builder.WriteString("}")
			} else {
				builder.WriteString("{}")
			}
		default:
			builder.WriteString(`""`)
		}
		if i == len(ptys)-1 {
			builder.WriteString("\n")
		} else {
			builder.WriteString(",\n")
		}
	}
}
