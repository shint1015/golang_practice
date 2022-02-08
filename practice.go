package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//基本的なログファイルの設定は以下のような感じで良い
func LoggingSetting(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSetting("test.log")
	_, err := os.Open("gdemnslkmw")
	if err != nil {
		log.Fatalln("exit", err)
	}

	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	//処理がそこで終了する
	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error")
	fmt.Println("ok")
}
