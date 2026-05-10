package main

import "fmt"

func main() {
	x := 10
	p := &x  // p é *int
	pp := &p // pp é **int

	**pp++
	fmt.Print(x)
}