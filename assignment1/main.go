package main

import (
	"fmt"

	"github.com/CoffeeSi/golangAITU/assignment1/Library"
	"github.com/CoffeeSi/golangAITU/assignment1/Shapes"
  "github.com/CoffeeSi/golangAITU/assignment1/Company"
  "github.com/CoffeeSi/golangAITU/assignment1/Bank"
  
)

func main() {
// 	library := new(Library.Library)
// 	library.ConsoleMenu()

	rectangle := Shapes.Rectangle{
		A: 4,
		B: 6,
	}
	circle := Shapes.Circle{
		R: 4,
	}
	square := Shapes.Square{
		A: 4,
	}
	triangle := Shapes.Triangle{
		A: 3,
		B: 4,
		C: 5,
	}

	shapes := []Shapes.Shape{rectangle, circle, square, triangle}

	for _, shape := range shapes {
		fmt.Printf("Area = %v; Perimetr = %v\n", shape.CalculateArea(), shape.CalculatePerimetr())
	}
  
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
  
  account := Bank.BankAccount{
		Name:          "Yevgeniy",
		AccountNumber: 12345678,
		Money:         0,
		Transactions:  []string{},
	}
	account.ConsoleMenu()
}
