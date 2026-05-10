package main

import "fmt"

func main() {
	num := 10

	fmt.Printf("num: %d - &num: %p\n", num, &num)

	var intPointer *int
	fmt.Println(*intPointer)
}