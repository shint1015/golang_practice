package main

import "fmt"

func main() {
	//Q1:出力結果を予測
	/*
		var i int = 10
		var p *int
		p = &i
		var j int
		j = *p
		fmt.Println(j)
	*/
	//Q2:出力結果を予測
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2
	//p1には、iのアドレスが入っているから
	//この時点で*p1=300
	//ここでp1のアドレスがp2に入る
	p2 = p1
	//fmt.Println(*p2)
	//*p2=300 + i = 300 = 600
	j = *p2 + i
	fmt.Println(j)
}
