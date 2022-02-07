package main

import "fmt"

//可変長引数　params ...int　（変数名 ...int）
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}

}

func main() {
	foo()
	foo(1, 2)
	foo(1, 2, 3, 4)

	s := []int{10, 20, 30}
	fmt.Println(s)
	//スライスの中を渡す時 => ...
	foo(s...)
}
