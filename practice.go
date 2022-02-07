package main

import "fmt"

func main() {

	//通常の配列作成
	l := []string{"python", "go", "php"}

	//普通に配列を回す時
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	//rangeを使って配列を回す
	//key【連番】,valueを回す時
	for i, v := range l {
		fmt.Println(i, v)
	}

	//key【連番】を回す時
	for v := range l {
		fmt.Println(v)
	}
	//valueを回す時
	for _, v := range l {
		fmt.Println(v)
	}

	//連想配列を作成
	m := map[string]int{"apple": 100, "banana": 150}

	//key,valueを回す時
	for k, v := range m {
		fmt.Println(k, v)
	}
	//valueだけを回す時
	for v := range m {
		fmt.Println(v)
	}
	//keyだけを回す時
	for _, k := range m {
		fmt.Println(k)
	}
}
