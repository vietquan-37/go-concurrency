package main

import (
	"fmt"
	"time"
)

func main() {
	//dowork := func(strings <-chan string) <-chan interface{} {
	//	completed := make(chan interface{})
	//	go func() {
	//		defer fmt.Println("dowork completed")
	//		defer close(completed)
	//		for s := range strings {
	//			fmt.Println("do work", s)
	//		}
	//	}()
	//	return completed
	//}
	//dowork(nil)
	//time.Sleep(time.Second*5)
	//fmt.Println("done")
	// memory leak if the long-lived software
	//solution
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("go routine exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					{
						fmt.Println(s)
					}
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	//Here we create another goroutine that will cancel the goroutine spawned in
	//doWork if more than one second passes.
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("go routine cancelled")
		close(done)
	}()
	<-terminated
	
	fmt.Println("done")
}
