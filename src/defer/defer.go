package main

import "fmt"

func foo() {
	defer fmt.Println("END!")
	fmt.Println("hey")
}

func main() {
	foo()
}
