package main

import "fmt"

func main() {

	t, f := true, false

	//%T 方表示
	//%v 値のデフォルトのフォーマットでの表現を出力
	//%t bool型を厳密に表したい場合
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false, "\n")

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false, "\n")

	fmt.Println(!true)
	fmt.Println(!false)
}
