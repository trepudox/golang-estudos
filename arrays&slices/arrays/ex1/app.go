package main

import "fmt"

func alteraArray(arr [2]int) {
	arr[0] = 99
}

func alteraArrayDeVdd(arr *[2]int) {
	arr[0] = 22
}

func main() {
	/*
	mesma coisa, os arrays tbm sao valores na memoria e é isso
	pra alterar o que está dentro do array precisamos passar o ponteiro dele
	*/
	arr := [2]int{0, 1}

	fmt.Printf("Valor inicial: %v\n", arr)
	
	alteraArray(arr)
	fmt.Printf("Valor alteraArray: %v\n", arr)

	alteraArrayDeVdd(&arr)
	fmt.Printf("Valor alteraArrayDeVdd: %v\n", arr)

}