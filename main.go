package main

import (
	"fmt"
	"sort"
)

//sort

func main() {
	i := []int{2, 1, 5, 9, 7, 3}
	s := []string{"A", "S", "D"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	//整数のソート
	sort.Ints(i)
	//文字列のソート
	sort.Strings(s)
	//struct?連想配列?の文字列ソート
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	//struct?連想配列?の数値ソート
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })

	fmt.Println(i, s, p)
}
