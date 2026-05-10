package main

import (
	"fmt"
	"time"
)

func main() {
	inicio := time.Now().UnixMilli()
	
	fmt.Println(fibonacci(50))
	
	fim := time.Now().UnixMilli()

	fmt.Println(inicio)
	fmt.Printf("tempo: %d", (fim - inicio))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}