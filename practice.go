package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	//if文を短縮して書くことができる
	//ただ、短縮したif文で定義した、変数は、if文以外で使用できない
	if result2 := by2(10); result2 == "ok" {
		fmt.Println(result2)
		fmt.Println("great 2")
	}
	//fmt.Println(result2)

	/*
		num := 6
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	x, y := 12, 13

	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}
