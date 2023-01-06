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
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(pi PrintInfo) {
	fmt.Println(pi.getMessage())
}

func main() {
	tEmployee := TemporaryEmployee{}
	ftEmployee := FullTimeEmployee{}

	// Polymorphic behaviour
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
