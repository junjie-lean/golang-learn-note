package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	fmt.Printf("arr1[0] = %d \n", arr1[0])
	arrConsole()
}

func arrConsole() {
	var arr2 [10]int
	for i := 0; i < len(arr2); i++ {
		if i == 0 {
			arr2[i] = 2
			continue
		}
		arr2[i] = arr2[i-1] * i
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Array at index of %d is %d\n", i, arr2[i])
	}
}
