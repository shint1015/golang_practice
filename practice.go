package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./practice.go")
	if err != nil {
		log.Fatal("error")
	}
	defer file.Close()
	data := make([]byte, 100)
	//変数の一個目だけイニシャライズすれば、大丈夫で
	//変数の2個目は、イニシャライズの対象にはならず、上書きされていく。
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("err")
	}
}
