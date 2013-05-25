package main

import "fmt"

//传入的是一个切片数组
func modify(arr []int) {
	arr[0] = 10
	fmt.Println("in func arr[0] = ", arr[0])
}

func main() {
	arr := []int {1, 2, 3, 4, 5}
	arr2 := []int {0}
	modify(arr)
	arr2 = append(arr, 90, 45, 90, 45, 33, 223, 4)

	fmt.Println("arr[] = ", arr)
	fmt.Println("arr2 = ", arr2)
}
