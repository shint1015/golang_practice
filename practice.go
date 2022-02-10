package main

import "fmt"

type Vertex struct {
	x, y int
}

func Area(v Vertex) int {
	return v.x * v.y
}

// 値レシーバー
func (v Vertex) Area() int {
	return v.x * v.y
}

// ポインタレシーバー
//　参照して、structの中身に変更を加える
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

//Golangのパターン
//newでストラクタを作成することがよくあるパターン
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	//v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
