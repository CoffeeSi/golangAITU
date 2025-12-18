package main

import (
	"fmt"

	"github.com/CoffeeSi/golangAITU/assignment1/Bank"
	"github.com/CoffeeSi/golangAITU/assignment1/Company"
	"github.com/CoffeeSi/golangAITU/assignment1/Library"
	"github.com/CoffeeSi/golangAITU/assignment1/Shapes"
)

func UseLibrarySystem() {
	library := new(Library.Library)
	library.ConsoleMenu()
}

func UseShapes() {
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
}

func UseCompanySystem() {
	company := Company.Company{}
	employeeYevgeniy := Company.PartTimeEmployee{
		ID:        1,
		FirstName: "Yevgeniy",
		LastName:  "Averyanov",
		Position:  "Junior",
		Salary:    300000,
		Shifts:    3,
	}
	employeeArlan := Company.FullTimeEmployee{
		ID:        2,
		FirstName: "Arlan",
		LastName:  "Tursyn",
		Position:  "Middle",
		Salary:    900000,
	}

	company.AddEmployee(employeeYevgeniy)
	company.AddEmployee(employeeArlan)
	company.ListEmployees()
}

func UseBankSystem() {
	account := Bank.BankAccount{
		Name:          "Yevgeniy",
		AccountNumber: 12345678,
		Money:         0,
		Transactions:  []string{},
	}
	account.ConsoleMenu()
}

func main() {
	fmt.Println("- Choose exercise:")
	fmt.Println(" 1. Library Management System")
	fmt.Println(" 2. Shapes & Interfaces")
	fmt.Println(" 3. Employee Management System")
	fmt.Println(" 4. Bank Account Simulation")
	fmt.Println(" 0. Exit")
	fmt.Print(": ")
	var exercise int
	fmt.Scan(&exercise)
	switch exercise {
	case 1:
		UseLibrarySystem()
	case 2:
		UseShapes()
	case 3:
		UseCompanySystem()
	case 4:
		UseBankSystem()
	case 0:
		return
	default:
		fmt.Println("Incorrect exercise!")
	}

}
