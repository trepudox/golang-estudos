package main

import (
	"errors"
	"fmt"
)

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

// método da struct - desconto em porcentagem
func (p *Produto) AplicarDesconto(desconto float64) error {
	if (desconto > 99 || desconto <= 0) {
		// return errors.New(fmt.Sprintf("Não é possivel aplicar um desconto de %.0f no produto '%s'", desconto, p.Nome))
		return fmt.Errorf("Não é possivel aplicar um desconto de %.0f no produto '%s'", desconto, p.Nome)
	}

	p.Preco = p.Preco * ((100 - desconto) / 100)
	return nil
}

// métodos fora da struct
func Comprar(p *Produto, quantidade int) {
	p.Quantidade += quantidade
}

func Vender(p *Produto, quantidade int) error {
	if quantidade > p.Quantidade {
		return errors.New("Estoque insuficiente")
	}

	p.Quantidade -= quantidade
	return nil
}

func main() {
	produto1 := Produto{Nome: "Garrafa de agua", Preco: 2.5, Quantidade: 10}
	fmt.Printf("produto1: %+v\n", produto1)
	produto1.AplicarDesconto(10)
	fmt.Printf("produto1: %+v\n", produto1)

	err := produto1.AplicarDesconto(1000)
	tratarErro(err)

	Comprar(&produto1, 10)
	fmt.Printf("produto1: %+v\n", produto1)
	Vender(&produto1, 2)
	fmt.Printf("produto1: %+v\n", produto1)

	err = Vender(&produto1, 1 << 63 - 1)
	tratarErro(err)
}

func tratarErro(err error) {
	if err != nil {
		fmt.Println("deu ruim: ", err)
	} else {
		fmt.Println("suave")
	}
}
