package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("aaaaaa")

	fmt.Println("HELLO WORLD")
	fmt.Println("HELLO " + "WORLD")
	fmt.Println(string("HELLO WORLD"[0]))

	var s string = "hello world"

	fmt.Println(strings.Replace(s, "h", "x", 1))
	//置換処理
	//strings.Replace
	s = strings.Replace(s, "h", "x", 1)
	fmt.Println(s)
	s = "hello world"
	fmt.Println(strings.Contains(s, "hello"))

	//バッククォートは、改行など勝手にエスケープしてくれる
	fmt.Println(`test
				test
test`)
	//よく聞くのは、バックスラッシュでの対応
	fmt.Println("\"")
	fmt.Println(`"`)
}
