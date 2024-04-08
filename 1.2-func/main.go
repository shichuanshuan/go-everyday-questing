package main

import (
	"fmt"
)

// 下面这段代码输出什么以及原因？

func hello() []string {
	return nil
}

func main() {
	h := hello // 将函数赋值给 h，而不是将返回值赋值给 h
	// h := hello() 这才是将返回值赋值给 h
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
