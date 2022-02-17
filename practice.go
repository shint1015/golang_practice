package main

import (
	"fmt"
	"sync"
)

//goroutine 並列処理

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//処理が終了するのを待つための処理
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	//goroutineのコマンドが終了しなくても、
	//通常の処理が終わってしまったら、処理が終了する
	normal("hello")
	//time.Sleep(2000 * time.Millisecond)
	wg.Wait()
}
