package main

import "fmt"

//goroutine 並列処理
//chanel goroutineと通常処理のデータの受け渡しで使う

func main() {
	//第二引数が、最大のチャネルの個数
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	//チャネルから値を取得している
	//x := <-ch
	//fmt.Println(x)
	//y := <-ch
	//fmt.Println(y)
	//この時点で、2個取ったので、チャネルが空っぽになっている
	//だから、再度チャネルに入れることができる

	//rangeでチャネルを回す時は、closeでチャネルを終了しないといけない
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}

}
