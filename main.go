package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [seu-nome]")
		return
	}

	name := os.Args[1]
	fmt.Printf("Hello, %s. Welcome!\n", name)
}
