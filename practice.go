package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	//第一引数？を外で定義して実行もできる
	sum := 1

	//for sum < 10(条件)
	for sum < 10 {
		//sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	//無限ループに入る
	//for {
	//	fmt.Println("hello")
	//}
}
