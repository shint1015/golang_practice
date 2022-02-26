package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

//ini
//configファイルから呼び出すことができる
//例 config.ini
// [web]
//	port=8080

//取得方法
//cfg, _ := ini.Load("config.ini")
//cfg.Section("web").Key("port").MustInt()

//MustString(default値)でカッコ内に値を入れることで、configファイルの値が空もしくはない場合にデフォルト値を参照する
//pkg詳細
//https://pkg.go.dev/gopkg.in/ini.v1

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").MustString(),
	}
}

func main() {
	fmt.Printf("%T %v \n", Config.Port, Config.Port)
	fmt.Printf("%T %v \n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v \n", Config.SQLDriver, Config.SQLDriver)
}
