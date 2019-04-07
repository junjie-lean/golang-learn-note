package main

const name string = "lean"

func main() {
	str := "lean"

	print(str, "===>", &str, "\n")
	print(name, "===>", &name, "\n")
}
