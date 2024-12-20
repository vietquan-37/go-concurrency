package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func printFiles(filenames string) {
	content, err := os.ReadFile(filenames)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
func main() {
	filesname := os.Args[1:]
	for _, filename := range filesname {
		go printFiles(filename)
	}
	time.Sleep(2 * time.Second)
}
