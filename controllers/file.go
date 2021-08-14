package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/services"
	"github.com/gorilla/mux"
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
		page = 1
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 50
	}
	files, totalCount, err := services.FetchFiles(search, sort, page, limit)
	jsonString, _ := json.Marshal(map[string]interface{}{"files": files, "totalCount": totalCount})
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	formFile, handler, err := r.FormFile("file")
	if err != nil {
		w.Header().Set("Content-Type", "applicatoin/json")
		jsonString, _ := json.Marshal(map[string]interface{}{"isSuccess": false, "message": err})
		w.Write(jsonString)
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

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	file, err := services.FetchFile(fileId)
	if err != nil {
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Title))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	bytes, _ := ioutil.ReadFile(strconv.Itoa(file.Id))
	w.Write(bytes)
}
