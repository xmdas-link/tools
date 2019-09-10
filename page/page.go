package page

import (
	"math"
)

type Page struct {
	Page    uint `json:"page"`
	Limit   uint `json:"limit"`
	PageMax uint `json:"page_max"`
	Total   uint `json:"total"`
}

func (p *Page) GetOffset() uint {
	return (p.Page - 1) * p.Limit
}

func (p *Page) CountPages(total uint) {
	if p.Page == 0 {
		p.Page = 1
	}

	if total == 0 {
		p.Page = 1
		p.PageMax = 1
	} else {
		p.PageMax = uint(math.Ceil(float64(total) / float64(p.Limit)))
		if p.Page > p.PageMax {
			p.Page = p.PageMax
		}
	}

	p.Total = total
}

func MakePage(page uint, limit uint, total uint) Page {

	var p = Page{
		Page:  page,
		Limit: limit,
	}

	p.CountPages(total)

	return p
}
