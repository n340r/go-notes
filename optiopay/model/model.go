package model

type Employee struct {
	Name      string
	Employees []*Employee
}

// ???
func (e *Employee) AddEmployees(employees ...*Employee) {
	e.Employees = append(e.Employees, employees...)
}
