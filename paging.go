package apui

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/chioas"
	"net/http"
)

func (b *Browser) writePagination(ctx aitch.ImperativeContext, req *http.Request, def *chioas.Path) {
	// todo does the endpoint support paging? (and how can we tell?)
}
