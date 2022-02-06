package main

import "fmt"

func main() {
	//アスキーコード
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("OK")
	fmt.Println(c)
	fmt.Println(string(c))

}
