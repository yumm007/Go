package main

import "fmt"

func main(){
	str1, str2 := "i am ", "world!"
	str3 := str1 + str2

	for i, c := range(str3) {
		fmt.Printf("str3[%d] = %c\n", i, c)
	}

}
