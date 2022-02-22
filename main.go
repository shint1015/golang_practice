package main

import (
	"fmt"
	"regexp"
)

//regexp
//正規表現

func main() {
	match, _ := regexp.MatchString("a([a-z0-9]+)e", "appl0e")
	fmt.Println(match)

	//MustCompile
	//正規表現の条件を使い回す時
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	//s := "view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

	//FindString
	//一致するときに、変数に値が入る
	fs := r2.FindString("/view/testaaa")
	fmt.Println(fs)
	fs2 := r2.FindString("/viewwww/test")
	fmt.Println(fs2)

	//FindStringSubmatch
	//条件と一致する時、スライスで値が帰ってくる
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

}
