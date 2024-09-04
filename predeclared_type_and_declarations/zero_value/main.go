package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var (
		i         int
		f         float64
		b         bool
		s         string
		p         *int
		sl        []int
		m         map[string]int
		ch        chan int
		st        struct{}
		fn        func()
		intf      interface{}
		structVar Person
	)

	fmt.Println("int:", i)
	fmt.Println("float64:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("*int:", p)
	fmt.Println("[]int:", sl)
	fmt.Println("map[string]int:", m)
	fmt.Println("chan int:", ch)
	fmt.Println("struct{}:", st)
	fmt.Println("func():", fn)
	fmt.Println("interface{}:", intf)
	fmt.Println("Person struct:", structVar)
}
