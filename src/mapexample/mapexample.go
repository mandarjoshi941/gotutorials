package main

import "fmt"

func main() {
	mymap := make(map[string]float64)
	mymap["Mandar"] = 26
	mymap["Anand"] = 25
	fmt.Println(mymap)
	for k, v := range mymap {
		fmt.Println(k, ":", v)
	}
	delete(mymap, "Mandar")
	fmt.Println(mymap)
}
