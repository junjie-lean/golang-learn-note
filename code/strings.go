package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaaaqweqasdnowejolean1123xxxx"
	fmt.Printf("%s的开头是aaaa吗？\n", str)
	fmt.Printf("%t\n", strings.HasPrefix(str, "aaaa"))

	fmt.Printf("_____________________________\n")

	fmt.Printf("%s的结尾是xxxx吗？\n", str)
	fmt.Printf("%t\n", strings.HasSuffix(str, "xxxx"))

	fmt.Printf("_____________________________\n")

	fmt.Printf("%s包含lean这个字符串吗？\n", str)
	fmt.Printf("%t\n", strings.Contains(str, "lean"))

	fmt.Printf("_____________________________\n")

}
