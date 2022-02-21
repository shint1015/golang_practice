package mylib

import "fmt"

//大文字スタート
//外部パッケージから呼び出すことができる
var Public string = "Public"

//小文字スタート
//このファイル内からしか呼び出すことができる
var private string = "private"

type Person struct {
	//名前
	Name string
	//年齢
	Age  int
}

func Say() {
	fmt.Println("Human")
	fmt.Println(private)
}
