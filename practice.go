package main

import "fmt"

func do(i interface{}) {
	//型がinterface型なのでintに直さないと計算ができない
	//タイプアサーション
	//interface型をint型に変換
	//ii := i.(int)
	////i = ii * 2
	//fmt.Println(i)
	//ii *= 2
	//fmt.Println(ii)
	//タイプアサーション
	//interface型をstring型に変換
	//ss := i.(string)
	//fmt.Println(ss + "!")

	//switch type
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I dont know %T\n", v)
	}

}

func main() {
	do(10)
	do("Mike")
	do(true)
}
