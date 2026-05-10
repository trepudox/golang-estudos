package main

import "fmt"

type Conta struct {
	Saldo float64
}

// Versão A
func (c Conta) DepositarA(valor float64) {
	c.Saldo += valor
}

// Versão B
func (c *Conta) DepositarB(valor float64) {
	c.Saldo += valor
}

func main() {
	minhaConta := Conta{Saldo: 100}

	minhaConta.DepositarA(50)
	fmt.Println("Saldo após A:", minhaConta.Saldo)

	minhaConta.DepositarB(50)
	fmt.Println("Saldo após B:", minhaConta.Saldo)
}