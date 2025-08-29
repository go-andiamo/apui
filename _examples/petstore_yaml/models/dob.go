package models

import (
	"github.com/go-andiamo/httperr"
	"strings"
	"time"
)

type DoB time.Time

func (r DoB) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(r).Format("2006-01-02") + `"`), nil
}

func (r *DoB) UnmarshalJSON(data []byte) error {
	if !strings.HasPrefix(string(data), `"`) || !strings.HasSuffix(string(data), `"`) {
		return httperr.NewUnprocessableEntityError("doB must be a string")
	}
	if dt, err := time.Parse("2006-01-02", string(data[1:len(data)-1])); err != nil {
		return httperr.NewUnprocessableEntityError("doB cannot be parsed")
	} else {
		*r = DoB(dt)
		return nil
	}
}
