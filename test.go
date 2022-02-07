package main

//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//

//func main() {
//	// URLのパターンとハンドラーを渡す
//	http.HandleFunc("/foo", fooHandlerFunc)
//
//	if err := http.ListenAndServe(":8888", nil); err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}
//

// `/foo` へリクエストがきたときに実行されるハンドラー
// http.ResponseWriterにコンテンツを書き込む。
// もちろんHTMLテンプレートを使ったりJSONを返すことも可能
//func fooHandlerFunc(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "/fooページです！！")
//	fmt.Println("testtest")
//	fmt.Fprintf(w, "testest")
//}


//package main
//
//import (
//"fmt"
//"net/http"
//)

//func handler(writer http.ResponseWriter, request *http.Request)	{
//	path := request.URL.Path[1:]
//
//	if path == "testest" {
//		fmt.Fprintf(writer, "Good Morning!! %v!", path)
//	} else {
//		fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
//	}
//
//}
//
//func main() {
//	http.HandleFunc("/", handler)
//	http.HandleFunc("/testest", handler)
//	http.ListenAndServe(":8080", nil)
//
//}