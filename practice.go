package main

import "fmt"

//func add(x int, y int) => (x, y int)に省略できる
//1個目の（）は引数
//2個目の（）は返り値
func add(x, y int) (int, int) {
	return x + y, x - y
}

//返す変数名を指定した時
/**
*	関数内で、resultを定義はしなくても良い、するとエラーになる
*	（return　変数名）で返さなくても、良い
*	何が返り値かをわかりやすくする
 */
func calc(price, item int) (result int) {
	result = price * item
	return
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	r, r2 := add(100, 200)
	fmt.Println(r)
	fmt.Println(r2)

	r3 := calc(1000, 20)
	fmt.Println(r3)


	//変数に関数を追加する
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)


	//並列化で「go」をつけて並列処理を実行することができるので覚えておいて損はない
	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}
