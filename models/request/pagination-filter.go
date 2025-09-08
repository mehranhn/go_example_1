package request

type PaginationFilter struct {
	Search string `query:"search" validate:"omitempty"`
	Page   uint   `query:"page" validate:"min=1"`
	Limit  uint   `query:"limit" validate:"min=1,max=100"`
}

func (filter *PaginationFilter) Offset() uint {
	return (filter.Page - 1) * filter.Limit
}
