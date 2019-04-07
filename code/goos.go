package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	goos := runtime.GOOS
	fmt.Printf("当前操作系统是：%s\n", goos, "\n")
	path := os.Getenv("path")
	fmt.Printf("环境变量值Path是：%s\n", path, "\n")
}
