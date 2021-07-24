package main

import (
	"fmt"
	"log"
	"net/http"
	//"strings"
	"github.com/HAYASAKA-Ryosuke/simple-file-storage/controllers"
)

func filesRouter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //オプションを解析します。デフォルトでは解析しません。

	if r.URL.Path == "/files/" && r.Method == "GET" {
		data := controllers.FetchFileList(r);
		fmt.Fprintf(w, data) //ここでwに入るものがクライアントに出力されます。
		return
	}
}

func main() {
//	env := "local"
//	database.DBInit(env)
//	db := database.GetDatabase()
//	defer db.Close()

	// ルーティング
	http.HandleFunc("/files/", filesRouter)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
