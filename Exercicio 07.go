package main

import "fmt"

type Conta struct {
	Titular string
	Saldo   float64
}

func main() {
	conta := Conta{
		Titular: "Ricardo",
		Saldo:   1300.0,
	}
	valorAdicionado(&conta)
	fmt.Print(conta)
}

func valorAdicionado(c *Conta) {
	c.Saldo = c.Saldo + 500.0
}
