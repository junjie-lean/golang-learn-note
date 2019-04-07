package main

import "fmt"

func main() {

	if sumresult, result := sum(1, 2); result {
		fmt.Println(sumresult)
	}
}

func sum(a int, b int) (int, bool) {

	return a + b, true
}
