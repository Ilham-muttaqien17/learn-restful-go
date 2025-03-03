package utils

import (
	"github.com/gofiber/fiber/v2"
)

type PaginationParams struct {
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	Offset        int    `json:"offset"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_dir"`
	IsDesc        bool
}

type MetaPagination struct {
	Limit     int   `json:"limit"`
	Page      int   `json:"page"`
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
}

type PaginationResponse[T interface{}] struct {
	Data T               `json:"data"`
	Meta *MetaPagination `json:"meta"`
}

type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

// Build pagination query
func BuildPaginationParams(ctx *fiber.Ctx) PaginationParams {
	query := ctx.Queries()
	limit := ConverStringToInt(query["limit"], 10)
	page := ConverStringToInt(query["page"], 1)
	offset := limit * (page - 1)
	sortBy := ToString(query["sort_by"], "id")
	sortDirection := ToString(query["sort_direction"], "asc")

	if sortDirection != string(Asc) && sortDirection != string(Desc) {
		sortDirection = "asc"
	}

	return PaginationParams{
		Limit:         limit,
		Page:          page,
		Offset:        offset,
		SortBy:        sortBy,
		SortDirection: sortDirection,
		IsDesc:        sortDirection == "desc",
	}
}
