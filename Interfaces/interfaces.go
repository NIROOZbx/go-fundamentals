package main

import "fmt"

type Shape interface{
	Area() float64
	Perimeter() float64
}




type Rectangle struct{
	Length float64
	Width float64
}


type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return 3.14*c.Radius*c.Radius
}

func (c Circle) Perimeter() float64{
	return 2*3.14*c.Radius
}

func (r Rectangle) Area() float64{
	return r.Length*r.Width
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.Length+r.Width)
}

func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("------")
}

func main() {
	rect := Rectangle{Length: 5, Width: 3}
	circle := Circle{Radius: 4}

	// Now your main function is much cleaner.
	fmt.Println("Rectangle Details:")
	printShapeDetails(rect)

	fmt.Println("Circle Details:")
	printShapeDetails(circle)
	var s Shape

	fmt.Println(s)
	
}