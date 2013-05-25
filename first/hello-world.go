package main

import "fmt"

func cumputer(value int , value2 float64)(result float64, err error) {
	result = float64(value) * value2
	return result, nil
}

func main() {
	res, _ := cumputer(3, 5.5)
	fmt.Println("Hello, world.")
	fmt.Println("res ", res)
	fmt.Println("const ", c0, c1, c2, c3, c4, c5)
}

const (
	c0	= iota
	c1
	c2
	c3 = 88
	c4 = iota
	c5
)
