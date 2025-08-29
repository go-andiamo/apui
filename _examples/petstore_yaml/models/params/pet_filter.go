package params

import (
	"github.com/go-andiamo/chioas"
	"github.com/go-andiamo/httperr"
	"net/http"
	"petstore_yaml/models"
	"strings"
	"time"
)

type PetFilter struct {
	Category string
	Name     string
	DoBEqual *time.Time
	DobFrom  *time.Time
	DobTo    *time.Time
	Order    []string
}

const (
	paramName     = "name"
	paramDoB      = "dob"
	paramCategory = "category"
	paramOrder    = "order"
)

func (f PetFilter) ToQueryParams() chioas.QueryParams {
	return chioas.QueryParams{
		{
			Name:        paramCategory,
			Description: "Filter by category",
		},
		{
			Name:        paramName,
			Description: "Search/filter by name",
		},
		{
			Name:        paramDoB,
			Description: "Filter by dob (date of birth)",
			Schema: &chioas.Schema{
				Type:   "string",
				Format: "date",
			},
		},
		{
			Name:        paramOrder,
			Description: "Order result by property",
		},
	}
}

func (f *PetFilter) Matches(pet *models.Pet) (matches bool) {
	matches = true
	if f.Name != "" {
		matches = strings.HasPrefix(strings.ToLower(pet.Name), f.Name)
	}
	if matches && f.Category != "" {
		matches = strings.EqualFold(f.Category, pet.Category.Name)
	}
	if matches && f.DoBEqual != nil {
		matches = f.DoBEqual.Format("2006-01-02") == time.Time(pet.DoB).Format("2006-01-02")
	}
	if matches && f.DobFrom != nil {
		matches = time.Time(pet.DoB).Format("2006-01-02") >= f.DobFrom.Format("2006-01-02")
	}
	if matches && f.DobTo != nil {
		matches = time.Time(pet.DoB).Format("2006-01-02") <= f.DobTo.Format("2006-01-02")
	}
	return matches
}

func PetFilterFromRequest(r *http.Request) (result *PetFilter, err error) {
	present := false
	result = &PetFilter{}
	if vals, ok := r.URL.Query()[paramCategory]; ok {
		if len(vals) == 1 {
			present = true
			result.Category = vals[0]
		} else {
			err = httperr.NewBadRequestErrorf("query paramater %q can only be specified once", paramCategory)
		}
	}
	if vals, ok := r.URL.Query()[paramName]; ok {
		if len(vals) == 1 {
			present = true
			result.Name = strings.ToLower(vals[0])
		} else {
			err = httperr.NewBadRequestErrorf("query paramater %q can only be specified once", paramName)
		}
	}
	if vals, ok := r.URL.Query()[paramDoB]; ok {
		if len(vals) == 1 {
			present = true
			if dt, e := time.Parse("2006-01-02", vals[0]); e == nil {
				result.DoBEqual = &dt
			} else {
				err = httperr.NewBadRequestError("invalid date of birth")
			}
		} else if len(vals) == 2 {
			present = true
			if dt, e := time.Parse("2006-01-02", vals[0]); e == nil {
				result.DobFrom = &dt
			} else {
				err = httperr.NewBadRequestError("invalid date of birth")
			}
			if dt, e := time.Parse("2006-01-02", vals[1]); e == nil {
				result.DobTo = &dt
			} else {
				err = httperr.NewBadRequestError("invalid date of birth")
			}
		} else {
			err = httperr.NewBadRequestErrorf("query paramater %q can only be specified once or twice", paramDoB)
		}
	}
	if vals, ok := r.URL.Query()[paramOrder]; ok {
		present = true
		for _, val := range vals {
			for _, order := range strings.Split(val, ",") {
				result.Order = append(result.Order, strings.ToLower(order))
			}
		}
	}
	if !present || err != nil {
		result = nil
	}
	return result, err
}
