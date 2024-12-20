package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playerTotal := 4
	for i := 0; i < 4; i++ {
		go PlayHandler(cond, &playerTotal, i)
		time.Sleep(time.Second * 1)

	}
}
func PlayHandler(cond *sync.Cond, playerTotal *int, i int) {
	cond.L.Lock()
	fmt.Println("player connected: ", i)
	*playerTotal--
	if *playerTotal == 0 {
		cond.Broadcast()
	}
	for *playerTotal > 0 {
		fmt.Println(i, "waiting for more player")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("game started all player connected")
}
