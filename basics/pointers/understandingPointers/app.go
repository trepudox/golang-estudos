package main

import "fmt"

type Funcionario struct {
	Nome    string
	Salario float64
}

/*
 EM RESUMO:
	&var - retorna o endereço de memória (ponteiro) de uma variável
	*ptr - retorna o valor de um ponteiro
	*int - é um TIPO que armazena ponteiros
*/

// Sem ponteiro: recebe uma CÓPIA. O original não muda.
func aumentarSemPonteiro(f Funcionario) {
	fmt.Printf("ponteiro do struct recebido: %p\n", &f)
	f.Salario += 1000
}

// Com ponteiro: recebe o ENDEREÇO. Altera o original.
func aumentarComPonteiro(f *Funcionario) {
	fmt.Printf("ponteiro do struct recebido: %p\n", f)
	f.Salario += 1000
}

func main() {
	func1 := Funcionario{"João", 3000}
	fmt.Printf("ponteiro original do struct: %p\n\n", &func1)

	aumentarSemPonteiro(func1)
	fmt.Println("Sem ponteiro:", func1.Salario) // Continua 3000

	fmt.Println()

	aumentarComPonteiro(&func1)
	fmt.Println("Com ponteiro:", func1.Salario) // Agora é 4000!
}
