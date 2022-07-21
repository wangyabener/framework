package models

type PageInfo struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
}
