package main

import (
	"fmt"
)

func main() {
	fmt.Println("Convert an integer to a floating-point number 🔢➡️")
	intNumber := 42
	floatNumber := float64(intNumber)
	fmt.Printf("Converted to float: %f\n", floatNumber)
	fmt.Println("-------------------------------------")

	fmt.Println("Convert a floating-point number to an integer")
	floatNumber2 := 3.14
	intNumber2 := int(floatNumber2)
	fmt.Printf("Converted to int: %d\n", intNumber2)
	fmt.Println("-------------------------------------")

	fmt.Println("Convert an integer to a string")
	intNumber3 := 123
	stringNumber := fmt.Sprintf("%d", intNumber3)
	fmt.Printf("Converted to string: %s\n", stringNumber)
	fmt.Println("-------------------------------------")

	fmt.Println("Convert a string to an integer")
	stringNumber2 := "456"
	intNumber4 := 0
	_, _ = fmt.Sscanf(stringNumber2, "%d", &intNumber4)
	fmt.Printf("Converted to int: %d\n", intNumber4)

	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)
	fmt.Println("-------------------------------------")

	fmt.Println("Convert a string to an byte[]")
	s := "SWIFT"
	fmt.Println("String value entered:", s)

	fmt.Println("After converting to Bytes value:")
	//convert string to bytes
	b := []byte(s)

	//display the slice of bytes using fmt.Println () function
	fmt.Println(b)
	fmt.Println("-------------------------------------")

	fmt.Println("Convert a string and Rune")
	convertingBetweenStringAndRune()
}

func convertingBetweenStringAndRune() {
	str := "Go"

	// Converting string to rune slice
	runes := []rune(str)
	for _, r := range runes {
		fmt.Printf("Rune: %c, Unicode code point: %U\n", r, r)
	}

	// Converting rune slice back to string
	newStr := string(runes)
	fmt.Println("New String:", newStr)
}
