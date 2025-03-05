package main

import (
	"fmt"
)

func main() {
	// 方式1
	for i := 0; i < 10; i++ {
		fmt.Println("方式1，第", i+1, "次循环")
	}

	// 方式2
	// b := 1
	// for b < 10 {
	// 	fmt.Println("方式2，第", b, "次循环")
	// }

	// 方式3，无限循环
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	// var started bool
	// var stopped atomic.Bool
	// for {
	// 	if !started {
	// 		started = true
	// 		go func() {
	// 			for {
	// 				select {
	// 				case <-ctx.Done():
	// 					fmt.Println("ctx done")
	// 					stopped.Store(true)
	// 					return
	// 				}
	// 			}
	// 		}()
	// 	}
	// 	fmt.Println("main")
	// 	if stopped.Load() {
	// 		fmt.Println("jinlaile")
	// 		break
	// 	}
	// }

	// 遍历数组
	var a [10]string
	a[0] = "Hello"
	for i := range a {
		fmt.Println("当前下标：", i)
	}
	for i, e := range a {
		fmt.Println("a[", i, "] = ", e)
	}

	// 遍历切片
	s := make([]string, 10)
	s[0] = "Hello"
	for i := range s {
		fmt.Println("当前下标：", i)
	}
	for i, e := range s {
		fmt.Println("s[", i, "] = ", e)
	}

	m := make(map[string]string)
	m["b"] = "Hello, b"
	m["a"] = "Hello, a"
	m["c"] = "Hello, c"
	for i := range m {
		fmt.Println("当前key：", i)
	}
	for k, v := range m {
		fmt.Println("m[", k, "] = ", v)
	}
}
