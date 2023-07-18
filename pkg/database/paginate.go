package database

type PageMeta struct {
	Current  int   `json:"current_page"`
	Last     int   `json:"last_page"`
	PageSize int   `json:"per_page"`
	Total    int64 `json:"total"`
}

type PaginationParams struct {
	Page     int `form:"page,default=1"`
	PageSize int `form:"per_page,default=10"`
}

func (p *PaginationParams) GetPage() int {
	// 默认第一页
	if p == nil || p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *PaginationParams) GetPageSize() int {
	// 默认每页 10 个
	if p == nil || p.PageSize == 0 {
		return 10
	}
	return p.PageSize
}
