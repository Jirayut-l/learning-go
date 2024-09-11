package main

import (
	slicingArrayMemory "Learning-Go/composite_types/array/slicing_array_memory"
	"fmt"
	"reflect"
)

func main() {

	//temp
	var arraysOriginal [3]int
	//let the compiler determine the length of an array based on the number of
	//elements provided during initialization.
	//Fixed Size: Arrays in Go are of a fixed size. The size is part of the array's type.
	var arraysEllipsis = [...]int{1, 2, 3, 4, 5} // can't add value
	fmt.Println(arraysOriginal)
	fmt.Println(arraysEllipsis)

	// array is fixed
	var numbers1 [5]int                 //creating an array of integers
	var numbers = [5]int{1, 2, 3, 4, 5} //initialize an array with values
	numbers3 := [5]int{1, 2, 3, 4, 5}   //shorthand
	fmt.Println(numbers1, numbers, numbers3)

	// Accessing Elements
	firstElement := numbers[0]  // Access the first element
	secondElement := numbers[1] // Access the second element
	numbers[2] = 10
	fmt.Println(numbers, firstElement, secondElement)

	// Iterating Over Arrays
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Length
	fmt.Println(len(numbers))

	//Passing Arrays to Functions
	printArray(numbers)

	// Zero values in arrays
	zeroValueArray()

	//Comparing Arrays
	comparingArrays()

	//Using reflect.DeepEqual()
	reflectDeepEqual()

	//Slicing array
	slicingArrayMemory.SlicingArrayMemory()
}

func printArray(arr [5]int) {
	for _, value := range arr {
		fmt.Println(value)
	}
}

func zeroValueArray() {
	// Creating an array of integers
	var numbers [5]int

	// Printing the array
	fmt.Println("Array before modification:", numbers)

	// Modifying the array
	numbers[2] = 42

	// Printing the modified array
	fmt.Println("Array after modification:", numbers)

	// Creating an array of strings
	var words [3]string

	// Printing the array
	fmt.Println("Array before modification:", words)

	// Modifying the array
	words[1] = "Hello"

	fmt.Println("Array after modification:", words)

}

func comparingArrays() {
	// Define two arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	// Compare arrays element by element
	equal := true
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			equal = false
			break
		}
	}

	// Print the result
	if equal {
		fmt.Println("Arrays are equal")
	} else {
		fmt.Println("Arrays are not equal")
	}
}

func reflectDeepEqual() {
	// Define two arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	// Compare arrays using DeepEqual
	if reflect.DeepEqual(arr1, arr2) {
		fmt.Println("Arrays are equal")
	} else {
		fmt.Println("Arrays are not equal")
	}
}
