/*
The main function for the calculator 
*/


package main 


import (
	"fmt"
	"log"

	"calculator/internal/calc"
)


func main() {


	a := 10
	b := 2

	fmt.Println("Calculator Example")
	fmt.Printf("%f + %f = %f\n", a, b, calc.Add(a, b))
	fmt.Printf("%f - %f = %f\n", a, b, calc.Subtract(a, b))
	fmt.Printf("%f * %f = %f\n", a, b, calc.Multiply(a, b))

	result, err := calc.Divide(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%f / %f = %f\n", a, b, result)



}