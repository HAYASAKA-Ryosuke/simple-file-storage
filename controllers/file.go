package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/services"
)

type JSONFile struct {
	id        string `json:"id"`
	title     string `json:"title"`
	createdAt string `json:"createdat"`
	updatedAt string `json:"updatedat"`
}

func FetchFileList(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(files)
	fmt.Println(totalCount)
	jsonString, _ := json.Marshal(map[string]interface{}{"files": files, "totalCount": totalCount})
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	formFile, handler, err := r.FormFile("file")
	if err != nil {
		return
	}

	defer formFile.Close()
	fmt.Println(handler.Filename)
	isSuccess, err := services.CreateFile(formFile, handler.Filename)
	if err != nil {
		return
	}
	jsonString, _ := json.Marshal(map[string]interface{}{"isSuccess": isSuccess})
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}
