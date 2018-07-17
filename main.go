package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a)
	var b, c = 4, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	switch {
	case a == 1:
		fmt.Println("a>1")
		break
	case b == 4:
		fmt.Println("a>1")
		break
	}
}
