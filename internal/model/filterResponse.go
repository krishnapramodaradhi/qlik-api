package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type FilterResponse struct {
	Region     string `json:"region"`
	Category   string `json:"category"`
	ScreenName string `json:"screenName"`
	UserName   string `json:"userName"`
	Role       string `json:"role"`
}

func (f *FilterResponse) Transform(filterName string, filterValues []string) *FilterResponse {
	switch filterName {
	case "region":
		f.Region = padFilterValues(filterValues)
	case "category":
		f.Category = padFilterValues(filterValues)
	case "screen_name":
		f.ScreenName = padFilterValues(filterValues)
	case "user_name":
		f.UserName = padFilterValues(filterValues)
	case "role":
		f.Role = padFilterValues(filterValues)
	default:
		log.Println("unknown filter", filterName)
	}
	return f
}

func (f FilterResponse) ScanRows(rows *sql.Rows) ([]FilterResponse, error) {
	filters := make([]FilterResponse, 0, 20)
	for rows.Next() {
		if err := rows.Scan(&f.Region, &f.Category, &f.ScreenName, &f.UserName, &f.Role); err != nil {
			return nil, err
		}
		filters = append(filters, f)
	}
	return filters, nil
}

func (f *FilterResponse) ParsedSearchQuery() string {
	query := "select * from utility_data"
	if f.Region != "" {
		query = query + fmt.Sprintf(" and region in (%v)", f.Region)
	}
	if f.Category != "" {
		query = query + fmt.Sprintf(" and category in (%v)", f.Category)
	}
	if f.ScreenName != "" {
		query = query + fmt.Sprintf(" and screen_name in (%v)", f.ScreenName)
	}
	if f.UserName != "" {
		query = query + fmt.Sprintf(" and user_name in (%v)", f.UserName)
	}
	if f.Role != "" {
		query = query + fmt.Sprintf(" and role in (%v)", f.Role)
	}
	return strings.Replace(query, "and", "where", 1)
}

func padFilterValues(filterValues []string) string {
	for i, v := range filterValues {
		filterValues[i] = "\"" + v + "\""
	}
	return strings.Join(filterValues, ",")
}
