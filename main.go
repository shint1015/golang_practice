package main

import (
	"fmt"
	"golang_project/mylib"
	"golang_project/mylib/under"
)

//他のパッケージからの呼び出しの際は、大文字スタートで宣言する必要がある

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	person := mylib.Person{Name: "taro", Age: 21}
	fmt.Println(person)

	fmt.Println(mylib.Public)

	//fmt.Println(mylib.private)

	under.Hello()
}
