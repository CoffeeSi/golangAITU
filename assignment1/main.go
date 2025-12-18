package main

import "github.com/CoffeeSi/golangAITU/assignment1/Company"

func main() {
	employee := Company.PartTimeEmployee{
		ID:        1,
		FirstName: "Yevgeniy",
		LastName:  "Averyanov",
		Position:  "Middle",
		Salary:    500000,
		Shifts:    3,
	}

	employee2 := Company.FullTimeEmployee{
		ID:        0,
		FirstName: "Arlan",
		LastName:  "Tursyn",
		Position:  "Senior",
		Salary:    1800000,
	}
	google := Company.Company{}
	google.AddEmployee(employee)
	google.AddEmployee(employee2)
	google.ListEmployees()
}
