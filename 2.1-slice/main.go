package main

import (  
    "fmt"
)

// 下面这段代码输出什么？
func main() {  
    a := [5]int{1, 2, 3, 4, 5}
    t := a[3:4:4]
    fmt.Println(t[0])
}