package main

var a string

func main() {
	a = "a"
	print(a) //a
	foo()
}

func foo() {
	a := "z"
	print(a) //z
	bar()
}

func bar() {
	print(a) //a
}
