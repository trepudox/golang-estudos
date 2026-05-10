package main

import (
	"fmt"
	"time"
)

func bancoDeDados(n int) {
	fmt.Printf("começando o bag %v\n", n)
	time.Sleep(2 * time.Second)
	fmt.Printf("terminando o bag%v\n", n)
}

func main() {
	for i := 0; i < 50; i++ {
		go bancoDeDados(i)
	}

	time.Sleep(5 * time.Second)
}