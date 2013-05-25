package main

import "fmt"

func cumputer(value int , value2 float64)(result float64, err error) {
	result = float64(value) * value2
	return result, nil
}

func main() {
	res, _ := cumputer(3, 5.5)
	fmt.Println("Hello, world. 你好！")
	fmt.Println("res ", res)
}
