package utils

import (
	"fmt"
	"math"
	"net/url"
)

type PaginationResponse struct {
	TotalItems  int         `json:"total_items"`
	TotalPages  int         `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
	Items       interface{} `json:"items"`
}

func PaginateResponse(data interface{}, totalItems, page, limit int) PaginationResponse {
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return PaginationResponse{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: page,
		Items:       data,
	}
}

func GetPageAndLimit(params url.Values) (int, int) {
	page := 1
	limit := 10
	pagination := true

	if val := params.Get("pagination"); val != "" {
		pagination = val == "true"
	}

	if pagination {
		if val := params.Get("page"); val != "" {
			fmt.Sscanf(val, "%d", &page)
		}
		if val := params.Get("limit"); val != "" {
			fmt.Sscanf(val, "%d", &limit)
		}

		if limit <= 0 {
			limit = 10
		}
	}

	return page, limit
}
