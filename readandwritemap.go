package main

import "time"

func main() {
	m := make(map[int]int)
	// 开启一段并发代码
	go func() {
		// 不停地对map进行写入
		for {
			m[1] = 1
		}
	}()
	// 开启一段并发代码
	go func() {
		// 不停地对map进行读取
		for {
			_ = m[1]
		}
	}()
	time.Sleep(10 * time.Second)
}
