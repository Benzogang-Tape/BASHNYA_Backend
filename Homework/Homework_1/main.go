package main

import "fmt"

func main() {
	fmt.Println("TEST")
	fmt.Print("ANOTHER TEST\n\n")
	var a, b float64
	fmt.Println("Enter two numbers to add them up: ")
	fmt.Scanf("%g\n%g", &a, &b)
	addition_result := test_function_add_numbers(a, b)
	fmt.Println(a, "+", b, "is", addition_result)
}

func test_function_add_numbers(a float64, b float64) float64 {
	return a + b
}
