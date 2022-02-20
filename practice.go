package main

import (
	"fmt"
	"sync"
	"time"
)

//goroutine 並列処理

//producer ,consumer
//並列処理を実行して、その結果のまとまりを、consumerで処理する

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		//fmt.Println("process", i * 1000)
		//wg.Done()
		//処理が正常に終了してから、Doneにする時はinner funcを使う手もある
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("##########################")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
	go consumer(ch, &wg)
	//処理が終了するのを待つ、以前使った
	wg.Wait()
	//chanelはcloseしないとconsumer内で次のrangeの値を待つため、closeしてあげないといけない。
	//close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
