package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	/*
		//deferは関数を実行した最後に実行される
		defer fmt.Println("world")

		foo()

		fmt.Println("hello")
	*/

	fmt.Println("run")
	//3,2,1の順で出力される => 再帰的な関数の実行みたいな？？
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	//defer:使い方の例
	file, _ := os.Open("./practice.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

}
