package main

import "fmt"

//Q1. 以下に、7と表示されるメソッドを作成してください。
//Q2. X is 3! Y is 4! と表示されるStringerを作成してください。

type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
	fmt.Println(v)
}
