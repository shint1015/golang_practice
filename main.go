package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//http

func main() {
	//URLからアクセス、取得ができる
	//resp, _ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//urlの中が、http://example.com/avssなどが入っていたとしても
	//http://example.comを取ってくる
	base, _ := url.Parse("http://example.com")
	//fmt.Println(base, err)
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)

	//postの場合は、領域を確保してあげて、渡したい値を渡す
	//req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))

	//ヘッダーに値を追加する
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	//パラメータを取得できる
	q := req.URL.Query()
	//パラメータを追加することができる
	q.Add("c", "3&%")
	q.Add("d", "")

	fmt.Println(q)
	fmt.Println(req.URL.RawQuery)
	//パラメータを変更したら、req.URL.RawQueryを更新してあげないといけない
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.RawQuery)

	var client *http.Client = &http.Client{}
	//client.Doで実際にアクセス
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
