package main

import (
	"fmt"
	"sync"
	"time"
)

//goroutine 並列処理
//同じ物に並列処理で変更を加えた際にエラーが起きることがある
//対応として、muxを使う

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	//変更を加える際は、対象への変更をロックできる
	c.mux.Lock()
	//変更を加えた後は、ロックを解除
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	//c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(500 * time.Millisecond)
	fmt.Println(c, c.Value("key"))
}
