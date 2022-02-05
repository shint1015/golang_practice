package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"

	//strconvのモジュールで文字列の数字を数値の型に返還などができる。
	//Atoiは、2つの返り値があるが, i,_ 「_」を使うことで、2つ目の返り値(今回の場合はエラー)は使わないとできる

	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "hello world"
	fmt.Println(string(h[1]))

}
