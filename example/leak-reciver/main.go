package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randStream := func(done <-chan interface{}) <-chan int {
		ch := make(chan int)
		go func() {
			//this never happen because the 3 loop block it
			defer fmt.Println("goroutine exited")
			defer close(ch)
			for {
				select {
				case <-done:
					return
				case ch <- rand.Intn(100):

				}
			}
		}()
		return ch
	}
	done := make(chan interface{})
	randCh := randStream(done)
	fmt.Println("3 random int")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d : %d\n ", i, <-randCh)
	}
	close(done)
	time.Sleep(1 * time.Second)
}
