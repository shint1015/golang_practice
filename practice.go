package main

import "fmt"

//stringer
//出力内容を変更、制御したい時に使うことができる

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	//return "My name is " + p.Name + "."
	//内容的には上記と同じ
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
