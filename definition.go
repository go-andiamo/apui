package apui

import (
	"github.com/go-andiamo/chioas"
	"net/http"
	"regexp"
	"strings"
)

func (b *Browser) findRequestDef(r *http.Request) (*chioas.Path, []*chioas.Path, []string) {
	paths := strings.Split(strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/"), "/")
	subs := make([]*chioas.Path, len(paths))
	if b.definition != nil {
		curr := b.definition.Paths
		var found *chioas.Path
		l := len(paths) - 1
		for i, path := range paths {
			if p, ok := curr["/"+path]; ok {
				if i == l {
					found = &p
					subs[i] = found
				} else {
					curr = p.Paths
					subs[i] = &p
				}
			} else if varPath := findPathVarMatch(path, curr); varPath != nil {
				if i == l {
					found = varPath
					subs[i] = found
				} else {
					curr = varPath.Paths
					subs[i] = varPath
				}
			}
		}
		return found, subs, paths
	}
	return nil, subs, paths
}

func findPathVarMatch(path string, curr chioas.Paths) (found *chioas.Path) {
	for k, rp := range curr {
		if strings.HasPrefix(k, "/{") && strings.HasSuffix(k, "}") {
			pth := strings.TrimSuffix(strings.TrimPrefix(k, `/{`), `}`)
			matches := !strings.Contains(pth, ":")
			if !matches {
				paramRx := strings.SplitN(pth, ":", 2)
				if rx, err := regexp.Compile(paramRx[1]); err == nil {
					matches = rx.MatchString(path)
				}
			}
			if matches {
				found = &rp
				break
			}
		}
	}
	return
}
