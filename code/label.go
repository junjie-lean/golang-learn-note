package main

import "fmt"

func main() {

LABEL1:
	// fmt.Println("好了 开始输出")
	for a := 0; a < 10; a++ {
		if a == 5 {
			continue LABEL1
		}
		fmt.Println(a)
	}

}
