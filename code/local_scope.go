package main

// import

var a = "g"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "0"
	print(a)
}
