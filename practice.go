package main

import "fmt"

//goroutine 並列処理

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	//chanは、処理が終了したら、closeする
	//closeしないと、rangeで回した時にエラーが起きる
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
