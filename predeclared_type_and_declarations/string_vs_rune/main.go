package main

import (
	"fmt"
)

func main() {

	fmt.Println("String")
	var str string = "Hello, 世界"
	fmt.Printf("String:%s", str)
	fmt.Println("-------------------------------------")

	fmt.Println("Rune")
	var r rune = '世'
	fmt.Printf("\nRune: %c, Unicode code point: %U", r, r)
	fmt.Println("-------------------------------------")

	//Instead a character is typically represented using a rune
	fmt.Println("Go does not have a specific 'char'")
	var char rune = 'A'
	fmt.Printf("Character: %c, Unicode code point: %U\n", char, char)

}
