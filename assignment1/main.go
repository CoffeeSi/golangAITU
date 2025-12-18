package main

import (
	"fmt"

	"github.com/CoffeeSi/golangAITU/assignment1/Shapes"
)

func main() {
	// library := new(Library.Library)
	// library.ConsoleMenu()

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
