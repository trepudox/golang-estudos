package main

import (
	"fmt"
	i "trepudox/desafio-interface/interfaces"
	s "trepudox/desafio-interface/structs"
)

func ImprimirFatura(p i.Pagavel) {
	fmt.Printf("Preco final: %.2f\n", p.PrecoTotal())
}

func main() {
	var livro i.Pagavel = &s.Livro{Nome: "Java Efetivo", Preco: 3999.99}
	var consultoria i.Pagavel = &s.Consultoria{Nome: "everis", Horas: 100, ValorHora: 75}

	ImprimirFatura(livro)
	ImprimirFatura(consultoria)
}
