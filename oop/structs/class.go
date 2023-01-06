package main

import "fmt"

// The structs are the go equivalent for classes
type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)

	e.id = 1
	e.name = "Bruno"
	fmt.Printf("%v", e)
}
