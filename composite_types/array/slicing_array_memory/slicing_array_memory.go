package slicing_array_memory

import (
	"fmt"
)

func SlicingArrayMemory() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
