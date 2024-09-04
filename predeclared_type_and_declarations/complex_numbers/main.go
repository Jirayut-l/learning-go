package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Declare complex numbers
	complexNum := 3 + 4i              // Complex number with inferred type (complex128)
	complexNum64 := complex64(1 + 2i) // Complex number explicitly typed as complex64

	// Basic operations with complex numbers
	sum := complexNum + complex128(complexNum64)
	difference := complexNum - complex128(complexNum64)
	product := complexNum * complex128(complexNum64)
	quotient := complexNum / complex128(complexNum64)

	// Display the results
	fmt.Println("Complex Number:", complexNum)
	fmt.Println("Complex Number (explicitly typed):", complexNum64)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Get the real and imaginary parts separately
	realPart := real(complexNum)
	imagPart := imag(complexNum)

	fmt.Println("Real Part:", realPart)
	fmt.Println("Imaginary Part:", imagPart)
	fmt.Println("Abs function:", cmplx.Abs(complexNum))

}
