package main

import "fmt"

func Calculator(a, b float64) {

	switch "division" {

	case "addition":
		fmt.Println(a + b)

	case "subtract":
		fmt.Println(a - b)

	case "division":
		fmt.Println(a / b)

	case "multiply":
		fmt.Println(a * b)

	}
}

func main() {

	Calculator(3, 5)

}
