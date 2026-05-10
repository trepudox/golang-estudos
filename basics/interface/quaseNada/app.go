package main

import "fmt"

type int64s interface {
	int64 | uint64
}

func printInts[T int64s](val T) {
	fmt.Printf("Type %T - Value: %v\n", val, val)
}

func main() {
	var maxInt int64 = 1 << 63 - 1
	var maxUint uint64 = 1 << 64 - 1

	printInts(maxInt)
	printInts(maxUint)

	val := 1 << 63 - 1
	fmt.Printf("Type %T - Value: %v\n", val, val)
}
