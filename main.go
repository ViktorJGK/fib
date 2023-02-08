package main

import (
	"fmt"

	"main.go/fib"
)

func main() {
	fmt.Println("Fantastiske Fibonacci!")
	result := fib.Fib(10)
	fmt.Println("Fib(10) =", result)
}
