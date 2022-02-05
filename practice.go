package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	//最初から配列内の個数を決めている場合は、新たに配列に追加することはできない
	//var b [2]int = [2]int{100, 200}
	//fmt.Println(b)
	//b = append(b, 300)

	//配列内の個数を決めてない場合は、新たに配列を追加することができる。
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
