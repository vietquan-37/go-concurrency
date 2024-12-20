package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	money := 100
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go Stingy(&money, cond)
	go Spendy(&money, cond)
	time.Sleep(2 * time.Second)
	fmt.Println("after: ", money)
	fmt.Println(time.Since(now))
}
func Stingy(money *int, mutex *sync.Cond) {
	for i := 0; i < 200000; i++ {
		mutex.L.Lock()
		for *money < 50 {
			mutex.Wait()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("negative money")
			os.Exit(1)
		}
		mutex.L.Unlock()
	}
	fmt.Println("Stingy done")
}
func Spendy(money *int, mutex *sync.Cond) {
	for i := 0; i < 1000000; i++ {

		mutex.L.Lock()
		*money += 10
		mutex.Signal()
		mutex.L.Unlock()

	}
	fmt.Println("Spendy done")
}
