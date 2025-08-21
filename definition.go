package apui

import (
	"github.com/go-andiamo/chioas"
	"net/http"
	"regexp"
	"strings"
)

func (b *Browser) findRequestDef(r *http.Request) (*chioas.Path, []*chioas.Path, []string, []string) {
	paths := strings.Split(strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/"), "/")
	subs := make([]*chioas.Path, len(paths))
	defPaths := make([]string, 0, len(paths))
	if b.definition != nil {
		curr := b.definition.Paths
		var found *chioas.Path
		l := len(paths) - 1
		for i, path := range paths {
			if p, ok := curr["/"+path]; ok {
				defPaths = append(defPaths, path)
				if i == l {
					found = &p
					subs[i] = found
				} else {
					curr = p.Paths
					subs[i] = &p
				}
			} else if varPath, varName := findPathVarMatch(path, curr); varPath != nil {
				defPaths = append(defPaths, "{"+varName+"}")
				if i == l {
					found = varPath
					subs[i] = found
				} else {
					curr = varPath.Paths
					subs[i] = varPath
				}
			}
		}
		return found, subs, paths, defPaths
	}
	return nil, subs, paths, defPaths
}

func findPathVarMatch(path string, curr chioas.Paths) (found *chioas.Path, varName string) {
	for k, rp := range curr {
		if strings.HasPrefix(k, "/{") && strings.HasSuffix(k, "}") {
			pth := strings.TrimSuffix(strings.TrimPrefix(k, `/{`), `}`)
			matches := !strings.Contains(pth, ":")
			varName = pth
			if !matches {
				paramRx := strings.SplitN(pth, ":", 2)
				varName = paramRx[0]
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
