package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}

	fmt.Println(m)
	fmt.Println(m["apple"])
	m["new"] = 500
	fmt.Println(m)

	//ない場合は、0が返ってくる
	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	//メモリ上に入れるmapがないので入れることができない
	//var m3 map[string]int
	//m3["pc"] = 5000
	//fmt.Println(m3)

	//varで宣言した時は、slice,map両方で「Nil」という値が入ってくる
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
