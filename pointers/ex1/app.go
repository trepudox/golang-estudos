package main

import "fmt"

func incrementa(n int) {
	n++
}

func incrementaReal(n *int) {
	*n++
}

func main() {
	x := 10
	incrementa(x)
	fmt.Printf("Ponto A: %d\n", x)
	// nao incrementa, pq estamos passando o valor de x, então é criado uma cópia na memória com o valor de x

	p := &x
	incrementaReal(p)
	fmt.Printf("Ponto B: %d\n", x)
	// incrementa, pq estamos passando o endereço onde está guardado o valor de x
}