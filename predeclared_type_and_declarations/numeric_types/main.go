package main

import (
	"fmt"
	"math"
)

func main() {
	// Integer Types
	var integer int = 42
	var signedInt int64 = -1234567890
	var unsignedInt uint = 789012345

	// Floating-point Types
	var floatNum float32 = 3.14
	var doubleNum float64 = 2.71828

	// Complex Types
	var complexNum complex64 = 1 + 2i
	var complexDoubleNum complex128 = 3 + 4i

	// Operations
	sum := integer + int(signedInt) // Need to explicitly convert to perform operations
	diff := doubleNum - float64(floatNum)
	product := complex128(complexNum) * complexDoubleNum
	quotient := float64(signedInt) / float64(integer)

	// Mathematical functions from the math package
	squareRoot := math.Sqrt(float64(integer))
	power := math.Pow(float64(integer), 2)

	fmt.Println("Numeric Types 🔢")
	fmt.Println("Integer:", integer)
	fmt.Println("Signed Integer:", signedInt)
	fmt.Println("Unsigned Integer:", unsignedInt)
	fmt.Println("Float:", floatNum)
	fmt.Println("Double:", doubleNum)
	fmt.Println("Complex:", complexNum)
	fmt.Println("Complex Double:", complexDoubleNum)
	fmt.Println("-------------------------------------")

	fmt.Println("\nOperations: ➕➖➗✖️")
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product (Complex):", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("-------------------------------------")

	fmt.Println("\nMathematical Functions: 🤖")
	fmt.Println("Square Root:", squareRoot)
	fmt.Println("Power of 2:", power)
	fmt.Println("-------------------------------------")
}
