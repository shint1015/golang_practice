package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "hogehoge"
}

func main() {
	//os := getOsName()
	//switch os {

	//短縮形で書くこともできる => 宣言した変数はswitch内でしか使えない
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("MAC!!")
	case "windows":
		fmt.Println("WINDOWS!!")
	default:
		fmt.Println("Default", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	//switchに最初に変数を渡さずにcase内で条件を記述する時
	switch {//変数なし
	case t.Hour() < 12://変数を含めた条件式
		fmt.Println("Good Morning!!")
	case t.Hour() >= 12:
		fmt.Println("Good Afternoon!!")

	}
}
