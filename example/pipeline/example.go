package main

import (
	"fmt"
	"time"
)

func main() {
	//multiple := func(values []int, multiplier int) []int {
	//	multiplierValues := make([]int, len(values))
	//	for i, v := range values {
	//		multiplierValues[i] = multiplier * v
	//	}
	//	return multiplierValues
	//}
	//add := func(values []int, additive int) []int {
	//	additiveValues := make([]int, len(values))
	//	for i, v := range values {
	//		additiveValues[i] = v + additive
	//	}
	//	return additiveValues
	//}
	//ints := []int{2, 3, 4, 5, 6, 7}
	//for _, v := range multiple(add(multiple(ints, 2), 1), 2) {
	//	fmt.Println(v)
	//}
	//this was simpler but it limit the ablity to scale and put a heavy work on for range
	//multiple := func(value int, multiplier int) int {
	//	return value * multiplier
	//}
	//add := func(value int, additive int) int {
	//	return value + additive
	//}
	//ints := []int{1, 2, 3, 4, 5}
	//for _, v := range ints {
	//	fmt.Println(add(multiple(v, 2), 1))
	//}
	//	better appoarch
	generator := func(done <-chan interface{}, intergers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range intergers {
				select {
				case <-done:
					return
				case intStream <- i:
					time.Sleep(time.Second)
				}
			}
		}()
		return intStream
	}
	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}
	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addStream <- i + additive:
				}
			}
		}()
		return addStream
	}
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
}
