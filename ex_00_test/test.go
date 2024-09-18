package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Demo struct {
	Date time.Time
}

func main() {
	demo := &Demo{
		Date: time.Now(),
	}
	jsonStr, _ := json.Marshal(demo)
	fmt.Println(string(jsonStr))
}
