package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c) // index of letter
		if cIndex >= 0 {
			//mutex.Lock()
			frequency[cIndex] += 1
			//mutex.Unlock()
		}
	}
	mutex.Unlock()
	fmt.Println("Completed:", url)
}

func main() {
	now := time.Now()
	var frequency = make([]int, 26)
	mutex := sync.Mutex{}
	for i := 1000; i <= 1010; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	time.Sleep(10 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	//++10s
	fmt.Println(time.Since(now).Seconds())
}

//race condition
