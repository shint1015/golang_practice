package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

//以下のような形で、渡すとキー名を変換できる
//`json:"name"`
//型を指定して渡す場合
//`json:"name, string(型名)"`
//jsonにする際に、非表示にする場合
//`json:"-"`
//空の場合、値を渡さない場合
//`json:"name, omitempty"`
//structにomitemptyを使う場合
//T *T（ポインタで渡す必要がある） `json:"T,omitempty"`

type Person struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
	T         *T       `json:"T,omitempty"`
}

//UnMarshalをカスタマイズする
//func (p *Person) UnmarshalJSON(b []byte) error {
//	type Person2 struct {
//		Name string
//		Age  int
//	}
//	var p2 Person2
//	err := json.Unmarshal(b, &p2)
//	if err != nil {
//		fmt.Println(err)
//	}
//	p.Name = p2.Name + "!"
//	p.Age = p2.Age * 10
//	return err
//}

//Marshalをカスタマイズする
//func (p Person) MarshalJSON() ([]byte, error) {
//	v, err := json.Marshal(&struct {
//		Name string
//	}{
//		Name: "Mr." + p.Name,
//	})
//
//	return v, err
//}

func main() {
	b := []byte(`{"name":"Tom","age":20, "nicknames":["a","b","c"]}`)
	var p Person
	//Unmarshal
	//ネットワークを入ってきたものを、そのままstructに入れてくれる
	//小文字でも大文字でも判定して、入れてくれる
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)
	//Marshal
	//jsonに変換する関数
	v, _ := json.Marshal(p)
	fmt.Println(string(v))

}
