package main

import "fmt"

type UserNotFound struct {
	UserName string
}

//エラーをカスタマイズすることができる関数
//ポインタじゃなくても、できるがポインタ渡しが主流
//ポインタで渡す理由は、
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found!! %v", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "mike"}
}

func main() {
	//ポインタで渡す理由は、別のエラーとして扱う為？？
	e1 := &UserNotFound{"mike"}
	e2 := &UserNotFound{"mike"}
	fmt.Println(e1 == e2)
	if err := myFunc(); err != nil {
		fmt.Println(err)
		//errorの種類によってハンドリングを分ける
		if err == e1 {

		}

		if err == e2 {

		}
	}
}
