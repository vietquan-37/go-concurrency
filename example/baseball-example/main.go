package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// listen to all event happens in the game
func matchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		//this modify the shared data match event
		mutex.Lock()
		*matchEvents = append(*matchEvents,
			"Match event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Appended match event")
	}
}

// this is just reader share data
func clientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		//this allow multiple goroutines can read data without blocking each other
		//,also prevent any write action while client reading
		mutex.RLock()
		//this dont modify the event this copy so we can use rlock
		allEvents := copyAllEvents(mEvents)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	for _, e := range *matchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}

func main() {
	mutex := sync.RWMutex{}
	var matchEvents = make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}
	go matchRecorder(&matchEvents, &mutex)
	start := time.Now()
	for j := 0; j < 5000; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)
}
