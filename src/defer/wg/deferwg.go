package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(msg string) {
	defer wg.Done()
	for idx := 0; idx < 5; idx++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
	wg.Add(1)
	go say("hi")
	wg.Wait()
}
