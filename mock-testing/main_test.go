package main

import "testing"

func TestGetFullTimeEmployeeByID(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{
						ID:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "John Doe",
					Age:  35,
					DNI:  "1",
				},
				Employee: Employee{
					ID:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeByID := GetEmployeeByID
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeByID(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeByID = originalGetEmployeeByID
	GetPersonByDNI = originalGetPersonByDNI
}
