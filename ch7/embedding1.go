package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

// func (m Manager) FindNewEmployees() []Employee {
// 	// do business logic
// }

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	fmt.Println("m.ID:", m.ID)
	fmt.Println("m.Description():", m.Description())
	fmt.Println()
	fmt.Println("m.Employee.Name:", m.Employee.Name)
	fmt.Println("m.Employee.ID:", m.Employee.ID)
}
