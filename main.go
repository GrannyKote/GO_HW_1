package main

import (
	"fmt"
	"os"
	"runtime"
)

func usernameFromEnv() string {
	if u := os.Getenv("USERNAME"); u != "" {
		return u
	}
	return os.Getenv("USER")
}

func main() {
	fmt.Println("Имя пользователя:", usernameFromEnv())

	fmt.Println("Аргументы CLI:")
	if len(os.Args) <= 1 {
		fmt.Println("  (нет аргументов)")
	} else {
		for i, arg := range os.Args[1:] {
			fmt.Printf("  [%d] %s\n", i+1, arg)
		}
	}

	fmt.Println("Версия Go:", runtime.Version())
}
