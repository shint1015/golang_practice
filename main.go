package main

//func main() {
//	s := []int{1, 2, 3, 4, 5}
//	fmt.Println(mylib.Average(s))
//
//	mylib.Say()
//	person := mylib.Person{Name: "taro", Age: 21}
//	fmt.Println(person)
//
//	fmt.Println(mylib.Public)
//
//	//fmt.Println(mylib.private)
//
//	under.Hello()
//}

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
