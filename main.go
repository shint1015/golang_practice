package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

//ioutil
//ネットワーク関係で使われる
//osパッケージと似ている
func main() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}
	//contentはbyteで返ってくる
	fmt.Println(string(content))

	//ファイル生成
	//if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	//	log.Fatalln(err)
	//}

	//bufferに関しても,readAllで読み込める
	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2))
}
