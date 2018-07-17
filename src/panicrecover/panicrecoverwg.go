package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("There There, have a tea: ", r)
	}
}
func say(msg string) {
	defer cleanup()
	for idx := 0; idx < 5; idx++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 300)
		if idx == 4 {
			panic("I don't like this")
		}
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
