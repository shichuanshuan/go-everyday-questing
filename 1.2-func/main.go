package main

import (
"fmt"
)

// 下面这段代码输出什么以及原因？

func hello() []string {  
    return nil
}

func main() {  
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}