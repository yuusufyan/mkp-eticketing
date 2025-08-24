package dto

import "auth-rbac/src/utils/enums"

type PaginateRequest struct {
	Page    int                  `query:"page" json:"page"`
	PerPage int                  `query:"perPage" json:"perPage"`
	Search  string               `query:"search" json:"search"`
	OrderBy string               `query:"orderBy" json:"orderBy"`
	SortBy  enums.OrderDirection `query:"sortBy" json:"sortBy"`
	Status  enums.StatusActive   `query:"status" json:"status"`
}

type ResponseDefault struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
}

type PaginationMeta struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// Value default di PaginationRequest
func (p *PaginateRequest) SetDefault() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 {
		p.PerPage = 10
	}
}
