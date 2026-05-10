package main

import "fmt"

func defaultFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forWithoutInitAndPostStatements() {
	sum := 1
	for ; sum < 100 ; {
		sum += sum
		fmt.Println(sum)
	}
}

func rangeFor() {
	for i := range 10 {
		fmt.Println(i)
	}
}


// nao existe While aqui, só for
func fakeWhile() {
	sum := 1
	for {
		fmt.Println(sum)
		sum += 1

		if sum > 50 {
			break
		}
	}
}

func main() {
	// defaultFor()
	// forWithoutInitAndPostStatements()
	// rangeFor()
	fakeWhile()
}
