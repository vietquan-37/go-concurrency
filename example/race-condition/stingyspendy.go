package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	money := 100
	mutex := sync.Mutex{}
	go Stingy(&money, &mutex)
	go Spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	fmt.Println("after: ", money)
}
func Stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money -= 1
		mutex.Unlock()
	}
	fmt.Println("Stingy done")
}
func Spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 1
		mutex.Unlock()
	}
	fmt.Println("Spendy done")
}
