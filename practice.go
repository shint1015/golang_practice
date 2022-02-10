package main

import "fmt"

type Vertex struct {
	//X,Y int　短縮した形でも表わせる
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
	//本当は以下のように書くのが正しいが、structの場合は勝手に変換してくれる
	//(*v).X = 1000
}

func main() {
	vv := Vertex{1, 2, "test"}
	changeVertex(vv)
	fmt.Println(vv)

	vv2 := &Vertex{1, 2, "test"}
	changeVertex2(vv2)
	fmt.Println(*vv2)
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	//yの値は指定していない場合は、0となる
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	//値だけの場合は、全部の値を入れないとエラーが起こる
	v3 := Vertex{1, 2, "test"}
	// ×　v3 := Vertex{1, 2}
	fmt.Println(v3)

	//v4, v5は同じ
	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	//newで宣言する時は、v7の宣言と同じ
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	//こっちの宣言の方が使われるかも
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

}
