package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  uint8
}

type Employee struct {
	ID       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM person WHERE dni = dni
	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	e, err := GetEmployeeByID(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
