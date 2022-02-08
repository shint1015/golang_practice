package main

import "fmt"

/*
panicは、どうすれば良いのかわからなくする要素があるので
なるべく使わない方が良い
recover()は、panicをキャッチして処理が終了しないようにすることができる
*/
func thirdPartyConnectDB() {

	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok?")
}
