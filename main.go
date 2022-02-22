package main

import (
	"context"
	"fmt"
	"time"
)

//context
//処理に時間がかかった時に、キャンセルできる

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	//contextの呼び出し
	ctx := context.Background()
	//context 終了条件の設定
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go longProcess(ctx, ch)

	//goroutineの終了する
	//cancel()

CTXLOOP:
	for {
		select {
		case <-ctx.Done()://ctxの終了に引っかかった時
			fmt.Println(ctx.Err())//errorの文言を表示
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("##################")

}
