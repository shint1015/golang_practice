package main

import "fmt"

func main() {
	//Q1 ____ float64 => int
	f := 1.11
	fmt.Println(f)
	//answer
	fmt.Println(int(f))
	//模範解答
	i := int(f)
	fmt.Println(i)

	//Q2 ____　以下の出力結果
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])
	//answer
	fmt.Println(5, 6)

	//Q3 ____ 出力結果がこれになるように：map[string]int map[Mike:20 Nancy:24 Messi:30]
	//answer
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
