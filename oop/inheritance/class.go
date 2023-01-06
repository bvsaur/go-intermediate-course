package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Bruno"
	ftEmployee.age = 23
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
}
