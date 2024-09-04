package main

import (
	"fmt"
)

func main() {

	isGoFun := true
	isJavaFun := false

	// Boolean operations
	isBothFun := isGoFun && isJavaFun   // Logical AND
	isEitherFun := isGoFun || isJavaFun // Logical OR
	isNotJavaFun := !isJavaFun          // Logical NOT

	// Conditional statements
	if isGoFun {
		fmt.Println("Go is fun!")
	} else {
		fmt.Println("Go is not fun.")
	}

	// Print the boolean values and results of operations
	fmt.Println("Is Go fun?", isGoFun)
	fmt.Println("Is Java fun?", isJavaFun)
	fmt.Println("Is both Go and Java fun?", isBothFun)
	fmt.Println("Is either Go or Java fun?", isEitherFun)
	fmt.Println("Is not Java fun?", isNotJavaFun)

}
