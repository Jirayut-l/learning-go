package main

import (
	"fmt"
)

func main() {
	x := 10
	var z int
	z = 10
	if x > 5 {
		x, y := 5, 20
		var z int
		z = 20
		fmt.Printf("x:%d,y:%d\n", x, y)
		fmt.Printf("z:%d\n", z)
	}
	fmt.Println(x)
	fmt.Println(z)
}
