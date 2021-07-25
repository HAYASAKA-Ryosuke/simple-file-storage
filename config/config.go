package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

//設定を指定するときに使う関数。起動時に読み込まれる
//envに指定する値はlocal,staging,production
func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// configの中身を返す
func GetConfig() *viper.Viper {
	return config
}
