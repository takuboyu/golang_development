package main

import "fmt"

func main() {
	var x int = 1
	//float型に型変換される
	xx := float64(x)
	// %fはfloat
	fmt.Printf("%T %v %f\n", xx, xx, xx)
}
