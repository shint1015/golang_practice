package main

import (
	"fmt"
	"time"
)

//goroutine 並列処理

func goroutine1(ch chan string) {
	for {
		ch <- "package from 1"
		time.Sleep(2 * time.Second)
	}
}

func goroutine2(ch chan int) {
	for {
		ch <- 200
		time.Sleep(2 * time.Second)
	}
}

func main() {
	//channelを分けて、routineを分けることができる
	c1 := make(chan string)
	c2 := make(chan int)
	go goroutine1(c1)
	go goroutine2(c2)


	//selectを使って分ける
	//書き方は以下のようなやり方でやる
	for {
		select {
		//<-c1の値が入ってきたら
		case msg1 := <-c1:
			fmt.Println(msg1)
		//<-c2の値が入ってきたら
		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}
}
