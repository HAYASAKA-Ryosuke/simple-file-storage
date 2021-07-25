package controllers

import (
	"net/http"
	"strconv"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/services"
)

func FetchFileList(r *http.Request) string {
	url, _ := r.URL.Parse(r.URL.String())
	query := url.Query()
	search := query.Get("search")
	sort := query.Get("sort")
	if sort == "" {
		sort = "Id"
	}
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err == nil {
		limit = 50
	}
	files, totalCount, err := services.FetchFiles(search, sort, page, limit)
	return "hello"
}
