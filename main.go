package main

import "fmt"

//iota
//自動採番

const (
	c1 = iota
	c2
	c3
)

//１を10回シフトする
const (
	_      = iota
	KB int = 1 << (10 * iota) //10
	MB                        //20
	GB                        //30
)

func main() {
	fmt.Println(c1, c2, c3) //0,1,2
	fmt.Println(KB, MB, GB)
}
