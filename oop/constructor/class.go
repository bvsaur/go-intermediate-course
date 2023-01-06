package main

import "fmt"

type Employee struct {
	id       int
	name     string
	isIntern bool
}

func NewEmployee(id int, name string, isIntern bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		isIntern: isIntern,
	}
}

func main() {
	/*
		There are some ways to create an instance an emulate the constructor behaviour
	*/

	// 1
	e1 := Employee{}
	fmt.Printf("%v\n", e1)

	// 2
	e2 := Employee{id: 1, name: "Bruno", isIntern: false}
	fmt.Printf("%v\n", e2)

	// 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 3
	e3.name = "Alonso"
	e3.isIntern = true

	// 4
	e4 := NewEmployee(4, "Alisson", false)
	fmt.Printf("%v\n", e4)
}
