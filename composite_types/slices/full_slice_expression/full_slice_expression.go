package full_slice_expression

import (
	"fmt"
)

func FullSliceExpression() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("---------------------------")
	nums := []int{1, 2, 3, 4, 5}
	subSlice := nums[1:4:4]
	fmt.Println("Original slice:", nums)
	fmt.Println("Sub-slice:", subSlice)
	fmt.Println("Length of sub-slice:", len(subSlice))
	fmt.Println("Capacity of sub-slice:", cap(subSlice))
}
