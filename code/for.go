package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("this is value of i: %d,\n", i)
	}

	// whileture()
	test()
}

func test() {

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d;\n", v) //输出5个0，因为这个时候的v是初始化的默认值 0
		v = 5
	}

	for i := 10; i < 30; {
		fmt.Printf("%d;\n", i) //死循环 以为没有加i自增条件，并且下面的代码也不执行；其实可以在内部定义i的自增
		i++
	}

	s := ""
	for s != "aaaaaaa" {
		fmt.Println("Value of s: ", s)
		s = s + "a"
	}

	for i, j, k := 0, 5, "a"; i < 3 && j < 100 && k != "aaaaa"; i, j, k = i+1, j+1, k+"a" {
		fmt.Println("value of i,j,k= ", i, j, k)
	}
}

func whileture() {
	fmt.Print("\n")
	j := 0
	for {
		if j%2 == 0 {
			continue
		}

		if j >= 1000 {
			break
		}
		fmt.Printf("j == %d \n", j)
		j++
	}
}
