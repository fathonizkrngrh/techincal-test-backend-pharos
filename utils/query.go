package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func BuildConditionQuery(params url.Values, searchParams []string) (string, []interface{}) {
	var whereClauses []string
	var searchClauses []string
	var args []interface{}
	counter := 1

	if isActive := params.Get("is_active"); isActive != "" {
		isActiveBool, err := strconv.ParseInt(isActive, 10, 64)
		if err == nil {
			whereClauses = append(whereClauses, fmt.Sprintf("is_active = $%d", counter))
			args = append(args, isActiveBool)
			counter++
		}
	}

	for _, param := range searchParams {
		if search := params.Get("search"); search != "" {
			searchClauses = append(searchClauses, fmt.Sprintf("%s ILIKE $%d", param, counter))
			args = append(args, "%"+search+"%")
			counter++
		}
	}
	var whereQuery string
	if len(whereClauses) > 0 {
		whereQuery = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	if len(searchClauses) > 0 {
		if whereQuery == "" {
			whereQuery = " WHERE " + strings.Join(searchClauses, " OR ")
		} else {
			whereQuery += " AND (" + strings.Join(searchClauses, " OR ") + ")"
		}
	}

	return whereQuery, args
}

func BuildPaginationQuery(params url.Values) string {
	limit := 10
	offset := 0

	if pagination := params.Get("pagination"); pagination == "true" {

		valLimit := params.Get("limit")
		valPage := params.Get("page")

		if valLimit != "" {
			fmt.Sscanf(valLimit, "%d", &limit)
		}

		if valPage != "" {
			var page int
			fmt.Sscanf(valPage, "%d", &page)

			offset = (page - 1) * limit
		}

		return fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	}

	return ""
}

func BuildOrderQuery(params url.Values) string {
	var query string
	orderBy := params.Get("order_by")
	orderType := params.Get("order_type")
	if orderBy != "" {
		query = fmt.Sprintf(" ORDER BY %s %s", orderBy, orderType)
	} else {
		query = " ORDER BY created_at DESC"
	}

	return query
}
