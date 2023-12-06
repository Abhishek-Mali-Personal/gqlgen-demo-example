package models

import (
	"math"

	"gorm.io/gorm"
)

var (
	PageLimit int
	OrderBy   string
)

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) SetTotalPages() {
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.GetLimit())))
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = PageLimit
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) SetPage(page *int) {
	if page != nil && *page > 0 {
		p.Page = *page
	}
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id " + OrderBy
	}
	return p.Sort
}

func (p *Pagination) Paginate(sortBy string) func(db *gorm.DB) *gorm.DB {
	p.Sort = sortBy
	p.SetTotalPages()
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}

func (p *Pagination) OptionalPaginate(sortBy string, page *int,
	paginate *bool,
) func(db *gorm.DB) *gorm.DB {
	p.Sort = sortBy
	if paginate != nil && !*paginate {
		p.Limit = int(p.TotalRows)
	} else {
		p.Page = *page
	}
	p.SetTotalPages()
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}

func Paginate(sortBy string, page *int,
	paginate *bool, value interface{}, pagination *Pagination, db *gorm.DB,
) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	pagination.Sort = sortBy
	if paginate != nil && !*paginate {
		pagination.Limit = int(totalRows)
	} else {
		pagination.Page = *page
	}
	pagination.SetTotalPages()
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func SetAfterQueryPagination(sortBy string, data ResultUnion, totalRows int64) *Pagination {
	paginate := new(Pagination)
	paginate.Sort = sortBy
	paginate.TotalRows = totalRows
	paginate.GetPage()
	paginate.SetTotalPages()
	paginate.Rows = data
	paginate.GetLimit()
	return paginate
}
