package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func multiplybyFive(c chan int, num int) {
	defer wg.Done()
	c <- num * 5
}

func main() {
	multiplyChannel := make(chan int, 10)
	for idx := 1; idx <= 10; idx++ {
		wg.Add(1)
		go multiplybyFive(multiplyChannel, idx)
	}
	wg.Wait()
	close(multiplyChannel)
	for value := range multiplyChannel {
		fmt.Println(value)
	}
}
