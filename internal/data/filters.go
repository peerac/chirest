package data

import (
	"math"
	"peerac/go-chi-rest-example/internal/validator"
)

type Filters struct {
	Page          int
	PageSize      int
	Sort          string
	Order         string
	SortSafeList  []string
	OrderSafeList []string
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f Filters) limit() int {
	return f.PageSize
}

type PageMeta struct {
	Current int   `json:"current,omitempty"`
	Size    int   `json:"size,omitempty"`
	First   int   `json:"first,omitempty"`
	Last    int   `json:"last,omitempty"`
	Total   int64 `json:"total,omitempty"`
}

func calculatePageMeta(total int64, page, pageSize int) PageMeta {
	if total == 0 {
		return PageMeta{}
	}

	return PageMeta{
		Current: page,
		Size:    pageSize,
		First:   1,
		Last:    int(math.Ceil(float64(total) / float64(pageSize))),
		Total:   total,
	}
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater that zero")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
	v.Check(validator.In(f.Order, f.OrderSafeList...), "order", "invalid order type")
}
