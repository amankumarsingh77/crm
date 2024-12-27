package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math"
	"strconv"
)

type Pagination struct {
	Page    int    `json:"page"`
	Size    int    `json:"count"`
	OrderBy string `json:"order_by"`
}

const (
	defaultSize = 10
)

func (p *Pagination) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		p.Size = defaultSize
		return nil
	}
	size, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	p.Size = size
	return nil
}

func (p *Pagination) SetPage(pageQuery string) error {
	if pageQuery == "" {
		p.Page = 0
		return nil
	}
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	p.Page = page
	return nil
}

func (p *Pagination) SetOrderBy(orderQuery string) {
	p.OrderBy = orderQuery
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) GetSize() int {
	return p.Size
}

func (p *Pagination) GetOrderBy() string {
	return p.OrderBy
}

func (p *Pagination) GetOffset() int {
	if p.Page == 0 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *Pagination) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%v", p.Page, p.Size, p.OrderBy)
}

func GetPaginationFromCtx(ctx echo.Context) (*Pagination, error) {
	q := &Pagination{}
	if err := q.SetPage(ctx.QueryParam("page")); err != nil {
		return nil, err
	}
	if err := q.SetSize(ctx.QueryParam("count")); err != nil {
		return nil, err
	}
	q.SetOrderBy(ctx.QueryParam("orderBy"))
	return q, nil
}

func GetTotalPages(totalCount, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}

func GetHasMore(currPage, totalCount, pageSize int) bool {
	return currPage > totalCount/pageSize
}
