package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/krishnapramodaradhi/qlik-api/internal/model"
	"github.com/krishnapramodaradhi/qlik-api/internal/util"
)

func SearchFilters(w http.ResponseWriter, r *http.Request, db *sql.DB) (int, error) {
	searchParam := r.URL.Query()["f"]
	if len(searchParam) == 0 {
		return http.StatusBadRequest, errors.New("search param is required")
	}
	filters := strings.Split(searchParam[0], "::")
	f := new(model.FilterResponse)
	for _, filter := range filters {
		filterSlice := strings.Split(filter, ":")
		filterName := filterSlice[0]
		filterSlice = filterSlice[1:]
		f = f.Transform(filterName, filterSlice)
	}
	rows, err := db.Query(f.ParsedSearchQuery())
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utilities, err := new(model.Utility).ScanRows(rows)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, util.WriteJSON(w, http.StatusOK, map[string]any{"result": utilities})
}

func FetchFilters(w http.ResponseWriter, _ *http.Request, db *sql.DB) (int, error) {
	rows, err := db.Query("select distinct region, category, screen_name, user_name, role from utility_data")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	filters, err := new(model.FilterResponse).ScanRows(rows)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, util.WriteJSON(w, http.StatusOK, map[string]any{"result": filters})
}
