package sliceappendstorage

import (
	"fmt"
)

func SliceAppendStorage() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, "z")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
