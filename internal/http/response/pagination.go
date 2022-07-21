package response

import "framework/database"

type Pagination[T interface{}] struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	Total       int64 `json:"total"`
	Pages       int64 `json:"pages"`
	Data        []T   `json:"data"`
}

func PageResponse[T interface{}](page *database.Page[T]) *Pagination[T] {
	return &Pagination[T]{
		CurrentPage: page.CurrentPage,
		PageSize:    page.PageSize,
		Total:       page.Total,
		Pages:       page.Pages,
		Data:        page.Data,
	}
}
