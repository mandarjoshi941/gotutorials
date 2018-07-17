package main

import (
	"fmt"
)

func multiplybyFive(c chan int, num int) {
	c <- num * 5
}

func main() {
	multiplyChannel := make(chan int)
	go multiplybyFive(multiplyChannel, 4)
	go multiplybyFive(multiplyChannel, 7)

	v1 := <-multiplyChannel
	v2 := <-multiplyChannel
	fmt.Println(v1, v2)
}
