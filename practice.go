package main

import "fmt"

//goroutine 並列処理
//chanel goroutineと通常処理のデータの受け渡しで使う

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//chanelに値を渡す矢印
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c)
	x := <-c
	fmt.Println(x)
}
