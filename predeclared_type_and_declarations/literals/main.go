package main

import (
	"fmt"
)

//Types of Literals in Go
//1. Integer Literals
//2. Floating-Point Literals
//3. Complex Number Literals
//4. String Literals
//5. Rune Literals
//6. Boolean Literals

func main() {
	// Integer literals
	var a int = 42
	var b int = 0x2A
	var c int = 052
	var d int = 0b101010

	// Floating-point literals
	var e float64 = 3.14
	var f float64 = 1.23e4

	// Complex number literals
	var g complex128 = 1.0 + 2.0i

	// String literals
	var h string = "Hello, World!\n"
	var i string = `Hello,
World!`

	// Rune literals
	var j rune = 'A'
	var r rune = '世'
	var v rune = '🤖'

	// Boolean literals
	var k bool = true
	var l bool = false

	// Print the literals
	fmt.Println("Integer literals:", a, b, c, d)
	fmt.Println("Floating-point literals:", e, f)
	fmt.Println("Complex number literal:", g)
	fmt.Println("String literals:", h, i)
	fmt.Println("Rune literal:", j)
	fmt.Printf("Rune: %c %c, Unicode code point: %U %U\n", r, v, r, v)
	fmt.Println("Boolean literals:", k, l)

}
