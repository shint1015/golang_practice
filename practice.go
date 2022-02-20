package main

import "fmt"

//goroutine 並列処理
//Q1:間違えを探せ

//修正前
//func goroutine1(arr []string, ch chan int) {
//修正後
func goroutine1(arr []string, ch chan string) {
	sum := ""
	for _, v := range arr {
		sum += v
		ch <- sum
	}
	close(ch)
}

func main() {
	arr := []string{"test1!", "test2!", "test3!", "test4!"}
	//修正前
	//ch := make(chan int)
	//修正後
	ch := make(chan string)
	go goroutine1(arr, ch)
	for v := range ch {
		fmt.Println(v)
	}
}
