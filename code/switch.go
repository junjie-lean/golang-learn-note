package main

func main() {
	var val int = 30

	switch {
	case val < 30:
		print("小于三十")
	case val == 30:
		print("等于三十")
		fallthrough
	case val >= 30:
		print("大于等于三十")
	}

}
