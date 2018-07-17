package main

import "fmt"

func foo() {
	for idx := 1; idx <= 5; idx++ {
		defer fmt.Println(idx, "END!")
	}
}

func main() {
	foo()
}
