package database

import (
	"database/sql"
	"fmt"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/config"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// 設定ファイルから設定を読み取ってデータベースに接続する関数
// envには実行環境ごとに用意されたymlファイルの名前を指定すること
// ymlファイルはconfigディレクトリの直下に配置すること
func DBInit(env string) {
	config.Init(env)
	c := config.GetConfig()
	dbFileName := c.GetString("database.fileName")
	var err error
	DB, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err.Error())
	}
}

// データベースに接続するための変数を取得するための関数
// DBInit関数を先に実行してから使うこと
func GetDatabase() *sql.DB {
	return DB
}
