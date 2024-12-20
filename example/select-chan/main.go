package main

import (
	"fmt"
	"time"
)

func main() {

	//c := make(chan interface{})
	////go func() {
	////	time.Sleep(5 * time.Second)
	////	close(c)
	////}()
	////fmt.Println("Blocking on read")
	//select {
	//case <-c:
	//	fmt.Println("block: ", time.Since(start))
	//case <-time.After(time.Second):
	//	fmt.Println("timeout")
	//}
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	WorkCount := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		WorkCount++
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Final Work Count:", WorkCount)
}
