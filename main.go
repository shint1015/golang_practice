package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

//semaphore
//goroutineの同時実行数を制限できる
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {

	//goroutineの制限数を超えた場合、goroutineを終了させる時
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}

	//goroutineの制限数を超えた場合、goroutineを待機させる時
	//s.Acquire(ctx, 1)でロックする
	//if err := s.Acquire(ctx, 1); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)

	fmt.Println("Finish!!")
}
