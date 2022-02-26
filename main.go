package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

//talib
//株価の分析に使える
//Rsiは株の売買の動きのグラフ
//株価のインディケータを扱えるライブラリ
func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
