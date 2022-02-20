package mylib

import "testing"

//テストファイルを作成する時
//ファイル名 + _ + testとする
//テスト関数に関しては
// Test + "テストしたい関数名"

//terminal上でテストを実行する時（全部）
//go test ./...

//テストライブラリ
//ginkgo
//https://onsi.github.io/ginkgo/
//gomega
//https://onsi.github.io/gomega/
var Debug bool = true

func TestAverage(t *testing.T) {
	if Debug == false {
		t.Skip("skip Reason")
	}
	v := Average([]int{1, 2, 3, 4, 5, 6, 7})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
