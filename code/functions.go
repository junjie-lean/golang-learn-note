package main

import "fmt"

func main() {
	// fmt.Println("get right")
	//调用一个函数的是的时候，参数可以从另个一个函数里返回，但是参数的个数必须一致
	// lastResult := callFunc(sum(4, 5)) //4+5+1

	// addan := add(pr())
	// fmt.Printf("the last Result is %d,\n", lastResult);
	// fmt.Printf("the last Result is %d,\n", addan)

	minist := getMin(4, 2, 3, 1, 44)

	fmt.Printf("最小的数是%d", minist)

}
func callFunc(num int, res bool, a bool) int {
	c := num
	if res {
		c++
	} else {
		c--
	}
	return c
}

func sum(a, b int) (int, bool, bool) {
	c := a + b
	return c, true, false
}

func pr() (a int, b int) {
	return 1, 2
}

func add(a, b int) int {
	return a + b
}

func getMin(num ...int) int {
	if len(num) == 0 {
		return 0
	}
	min := num[0]
	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}
