package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("☑️ if and else")
	ifElseCondition()
	fmt.Println("-------------------------------------------")
	fmt.Println("✅ Ladder")
	ifLadderCondition()
	fmt.Println("-------------------------------------------")
	fmt.Println("🚩 Scoping a variable")
	scopingVariableAnIfStatement()
	fmt.Println("-------------------------------------------")
}

// if and else
func ifElseCondition() {
	n := rand.Intn(10)

	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func ifLadderCondition() {
	age := 25

	if age < 13 {
		fmt.Println("You are a child.")
	} else if age < 18 {
		fmt.Println("You are a teenager.")
	} else if age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}
}

func scopingVariableAnIfStatement() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}
