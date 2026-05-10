package main

import "fmt"

type MeuErro struct {
	Codigo   int
}

func (e *MeuErro) Error() string {
	return "deu bo"
}

func retornaErro() error {
	return &MeuErro{Codigo: 1}
}

func main() {
	err := retornaErro()
	if err != nil {
		if e, ok := err.(*MeuErro); ok {
			fmt.Printf("Erro - msg: %s, codigo: %d", e, e.Codigo)
		} else {
			fmt.Println("erro generico")
		}
	}
}
