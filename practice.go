package main

import "fmt"

//ポインタ
//アドレス、参照渡し
func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(*&*&n)
	fmt.Println(&n)
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)
}
