package main

import "fmt"

func IsEven(x int) bool {
	/* 
	codigo bunda e sem sentido mas só pra mostrar a capacidade do if
	eu posso criar uma variavel antes da verificacao booleana, 
	essa variavel entao fica disponivel dentro do escopo tanto do if quanto do else
	a sintaxe é essa:
	if <declaracao da variavel> ; <expressao booleana> {}
	*/
	if isEven := x % 2 == 0; isEven {
		fmt.Printf("%d é par ss\n", x)
		return isEven
	} else {
		fmt.Printf("%d é par nn\n", x)
		return isEven
	}
}

func main() {
	var num int;
	fmt.Print("Escolha um numero: ")
	fmt.Scan(&num)
	fmt.Println()
	fmt.Println(IsEven(num))
}