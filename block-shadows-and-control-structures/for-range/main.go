package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("The for-range for Statement")
	forRangeLoop()
	fmt.Println("-------------------------------------------")
	fmt.Println("Ignoring the slice index in a for-range loop")
	ignoringIndexValueOnly()
	fmt.Println("-------------------------------------------")
	fmt.Println("for-range loop Key Only")
	forRangeLoopKeyOnly()
	fmt.Println("-------------------------------------------")
	fmt.Println("Map Iteration Order Varies")
	mapIterationOrderVaries()
	fmt.Println("-------------------------------------------")

}

func forRangeLoop() {
	evenValues := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, v := range evenValues {
		fmt.Printf("index:%d,Value:%d\n", i, v)
	}
}

func ignoringIndexValueOnly() {
	evenValues := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenValues {
		fmt.Printf("Value:%d\n", v)
	}
}

func forRangeLoopKeyOnly() {
	//different initial map
	//uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	uniqueNames := make(map[string]bool)
	uniqueNames["Fred"] = true
	uniqueNames["Raul"] = true
	uniqueNames["Wilma"] = true

	for key := range uniqueNames {
		fmt.Printf("Key:%s\n", key)
	}
}

func mapIterationOrderVaries() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("iteration:", i)
		for k, v := range m {
			fmt.Printf("key:%s,value:%d\n", k, v)
		}
		fmt.Println("--------------------------")
	}
}
