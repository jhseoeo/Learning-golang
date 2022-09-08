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
	// Employee as an embedded field
	Employee // no name assigned to this filed. only type.
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return m.Reports
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)            // prints 12345
	fmt.Println(m.Description()) // prints Bob Bobson (12345)
}
