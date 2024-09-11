package main

import (
	confusingSlices "Learning-Go/composite_types/slices/confusing_slices"
	fullSliceExpression "Learning-Go/composite_types/slices/full_slice_expression"
	sliceappendstorage "Learning-Go/composite_types/slices/slice_append_storage"
	slicesharestorage "Learning-Go/composite_types/slices/slice_share_storage"
	slicingSlices "Learning-Go/composite_types/slices/sliceing_slices"
	"fmt"
)

func main() {

	//declaring
	var mySlice []int //zero value or null
	//declaring a slice with default values
	var mySliceData = []int{1, 2, 3, 4, 5}
	myShortSlice := []int{1, 2, 3, 4, 5} //shorthand
	fmt.Println(len(mySlice))
	fmt.Println(myShortSlice)
	fmt.Println(mySliceData)

	//slicing slices
	fmt.Println("slicing slices")
	slicingSlices.SlicingSlices()
	fmt.Println("----------------------------")

	//slice share storage
	fmt.Println("slice share storage")
	slicesharestorage.SliceShareStorage()
	fmt.Println("----------------------------")

	//slice append storage
	fmt.Println("slice append storage")
	sliceappendstorage.SliceAppendStorage()
	fmt.Println("----------------------------")

	//confusingSlices
	fmt.Println("confusingSlices")
	confusingSlices.ConfusingSlices()
	fmt.Println("----------------------------")

	//full
	fmt.Println("full slice expression")
	fullSliceExpression.FullSliceExpression()
	fmt.Println("----------------------------")
}
