package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(n)
	fmt.Println(n[2])
	//2<=x<4の間の値を表示 [3,4]
	fmt.Println(n[2:4])
	//x<2の間の値を表示 [1,2]
	fmt.Println(n[2:])
	//x>=2の間の値を表示 [3,4,5,6]
	fmt.Println(n[:2])
	n[2] = 100
	fmt.Println(n)

	//多次元配列は以下のように作成することができる。
	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)

	//appendで複数の値を入れることもできる
	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)

	board[1] = append(board[1], 100, 200, 300)
	fmt.Println(board)
	//多次元配列の新しい配列の追加
	board = append(board, []int{10, 11, 12})
	fmt.Println(board)
}
