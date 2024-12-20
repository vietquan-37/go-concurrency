package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permit int
	cond   *sync.Cond
}

func NewSemaphore(permit int) *Semaphore {
	return &Semaphore{
		permit: permit,
		cond:   sync.NewCond(&sync.Mutex{}),
	}
}
func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	for s.permit <= 0 {
		s.cond.Wait()
	}
	s.permit--
	s.cond.L.Unlock()
}
func (s *Semaphore) Release() {
	s.cond.L.Lock()
	s.permit++
	s.cond.Signal()
	s.cond.L.Unlock()
}
func main() {
	semaphore := NewSemaphore(0)
	for i := 0; i < 5000; i++ {
		go doWork(semaphore)
		fmt.Println("waiting for child routines")
		semaphore.Acquire()
		fmt.Println("child routines finshed")

	}
}
func doWork(semaphore *Semaphore) {
	fmt.Println("work started")
	fmt.Println("work finished")
	semaphore.Release()
}
