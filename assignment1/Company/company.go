package Company

import "fmt"

type Employee interface {
	GetDetails()
	GetId() uint64
}

type FullTimeEmployee struct {
	ID        uint64
	FirstName string
	LastName  string
	Position  string
	Salary    float64
}

type PartTimeEmployee struct {
	ID        uint64
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	Shifts    int
}

type Company struct {
	employees map[uint64]Employee
}

func (e FullTimeEmployee) GetDetails() {
	fmt.Printf("ID: %v, Name: %s %s, Position: %s, Salary: %v tg\n",
		e.ID, e.FirstName, e.LastName, e.Position, e.Salary)
}

func (e FullTimeEmployee) GetId() uint64 {
	return e.ID
}
func (e PartTimeEmployee) GetId() uint64 {
	return e.ID
}

func (e PartTimeEmployee) GetDetails() {
	fmt.Printf("ID: %v, Name: %s %s, Position: %s, Salary: %v tg, Shifts: %v\n",
		e.ID, e.FirstName, e.LastName, e.Position, e.Salary, e.Shifts)
}

func (company *Company) AddEmployee(employee Employee) {
	if company.employees == nil {
		company.employees = make(map[uint64]Employee)
	}

	var id uint64 = employee.GetId()

	if company.employees[id] != nil {
		fmt.Printf("Employee with ID=%v already exists\n", id)
		return
	}
	company.employees[id] = employee
	fmt.Printf("Employee with ID=%v successfully added!\n", id)
}

func (company Company) ListEmployees() {
	fmt.Println("List of employees:")
	for _, employee := range company.employees {
		employee.GetDetails()
	}
}
