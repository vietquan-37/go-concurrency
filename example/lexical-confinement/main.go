package main

import "fmt"

func main() {
	//write only
	producer := func() <-chan int {
		result := make(chan int, 5)
		go func() {
			defer close(result)
			for i := 0; i <= 5; i++ {
				result <- i
			}
		}()
		return result
	}
	//read only
	consumer := func(result <-chan int) {
		for n := range result {
			fmt.Println(n)
		}
		fmt.Println("Done receiving")
	}
	results := producer()
	consumer(results)
}
