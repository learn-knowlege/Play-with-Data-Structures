package main

import "fmt"

func main() {
	a := getArray(5)

	fmt.Println(a)
	fmt.Println(a.getCapacity(), a.getSize(), a.isEmpty())
}
