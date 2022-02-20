package main

import (
	"fmt"
	"time"
)

//goroutine 並列処理

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	//for breakを使う場合は,変数名：でスタートし
	//break 変数名でやればbreakすることができる
OuterBreak:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			//処理自体終了してします
			return
			//for文を終わらせるためには、
			break OuterBreak
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)

		}
	}
	fmt.Println("####################")
}
