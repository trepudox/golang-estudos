package main

import "fmt"

type Perfil struct {
	Nome string
}

func main() {
	// ! o problema: uma variavel de ponteiro sem valor, literalmente, null pointer exception (ou nil pointer pros golangers)
	var p *Perfil
	if p == nil {
		fmt.Println("tá nulo")
	}

	// a solução:
	p = new(Perfil)
	if p != nil {
		fmt.Println("não tá nulo!!")
		fmt.Println(p)
		fmt.Println(p.Nome)
	}

	p = nil

	// OU
	p = &Perfil{}
	if p != nil {
		fmt.Println("não tá nulo tb!!")
		fmt.Println(p)
		fmt.Println(p.Nome)
	}
}