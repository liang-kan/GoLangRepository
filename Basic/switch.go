package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("西华你好switch")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("你好坏")
	case "linux":
		fmt.Print("linux")
	default:
		fmt.Print("爱死你了")
	}
}
