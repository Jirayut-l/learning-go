package main

import (
	"fmt"
)

func main() {
	fmt.Println("🔴 Confusing Code")
	confusingCode()
	fmt.Println("-------------------------------------------")
	fmt.Println("🟢 Using continue to make code clearer")
	simplerFizzBuzz()
	fmt.Println("-------------------------------------------")
}

func confusingCode() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func simplerFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
