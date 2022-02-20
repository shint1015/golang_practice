package mylib

import "fmt"

///大文字スタート
//外部パッケージから呼び出すことができる
var Public string = "Public"

///小文字スタート
//外部パッケージから呼び出すことができる
var private string = "private"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human")
	fmt.Println(private)
}
