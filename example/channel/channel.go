package main

import "fmt"

func main() {

	////var receiveChan <-chan int
	////var sendChan chan<-int
	//stringChan := make(chan string)
	//go func() {
	//	//deadlock
	//	//if 0 != 1 {
	//	//	return
	//	//}
	//	stringChan <- "hello world"
	//
	//}()
	////close(stringChan)
	//fmt.Println(<-stringChan)
	chanOwner := func() <-chan int {
		c := make(chan int, 5)
		go func() {
			defer close(c)
			for i := 0; i <= 5; i++ {
				c <- i
			}

		}()
		return c
	}
	result := chanOwner()
	for v := range result {
		fmt.Println(v)
	}
	fmt.Println("Done")
}
