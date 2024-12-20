package main

import (
	"fmt"
	"time"
)

func countdown(count *int) {
	for *count > 0 {
		time.Sleep(time.Second)
		*count -= 1

	}
}
func main() {
	count := 5

	go countdown(&count)
	if count > 0 {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(count)
	}
}
