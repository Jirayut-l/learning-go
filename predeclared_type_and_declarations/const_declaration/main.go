package main

import "fmt"

const (
	Pi      float64 = 3.14159
	Version string  = "1.0.0"
)
const x int64 = 10
const z = 20 * 10

func main() {
	const y = "hello"
	fmt.Println("Pi:", Pi)
	fmt.Println("Version:", Version)
	fmt.Println("x:", x)
	fmt.Println("z:", z)
	fmt.Println("y:", y)
}
