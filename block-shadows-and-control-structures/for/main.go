package main

import "fmt"

func main() {
	fmt.Println("1️⃣ The Complete for Statement")
	completeFor()
	fmt.Println("-------------------------------------------")
	fmt.Println("2️⃣ The Condition-Only for Statement")
	conditionOnlyFor()
	fmt.Println("-------------------------------------------")
	fmt.Println("3️⃣ The infinite for Statement")
	infiniteLooping()
	fmt.Println("-------------------------------------------")
	fmt.Println("4️⃣ The for-range for Statement")
	forRangeLoop()
	fmt.Println("-------------------------------------------")
}

func completeFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func conditionOnlyFor() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func infiniteLooping() {
	i := 10
	for {
		fmt.Println("Hello World")
		i--
		//things to do in the loop
		if i <= 0 {
			break
		}
	}
	// Go has no equivalent of the do keyword
}

func forRangeLoop() {
	evenValues := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenValues {
		fmt.Printf("index:%d,Value:%d\n", i, v)
	}
}
