package main

import "fmt"

//goroutine 並列処理

//以下みたいな感じで、処理が終わったら
//次の処理みたいな感じで処理を実行していける
//goroutine1 => goroutine2 => goroutine3

func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first chan int, second chan int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

//chanelの値渡しの関係を引数で明示することもできる
func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func main() {

	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}
}
