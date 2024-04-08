package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	// marshaler 优先调用 MarshalJSON, 其次调用 MarshalText 方法
	// 这两个方法都不存在，则走正常的类型编码逻辑
	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}
