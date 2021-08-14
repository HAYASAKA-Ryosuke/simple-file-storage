package main

import (
	"log"
	"net/http"

	//"strings"
	"github.com/HAYASAKA-Ryosuke/simple-file-storage/controllers"
	"github.com/HAYASAKA-Ryosuke/simple-file-storage/database"
	"github.com/gorilla/mux"
)

func filesRouter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.URL.Path)
	log.Println(r.Method)

	if r.URL.Path == "/api/files/" && r.Method == "GET" {
		controllers.FetchFileList(w, r)
		return
	}
	if r.URL.Path == "/api/files/" && r.Method == "POST" {
		controllers.CreateFile(w, r)
		return
	}
}

func fileDownloadRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controllers.DownloadFile(w, r)
		return
	}
}

func main() {
	env := "local"
	database.DBInit(env)
	db := database.GetDatabase()
	defer db.Close()

	// ルーティング
	router := mux.NewRouter()
	router.HandleFunc("/api/files/{id:[0-9]+}/", fileDownloadRouter)
	router.HandleFunc("/api/files/", filesRouter)
	http.Handle("/", router)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
