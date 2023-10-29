package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	preco  float64
}

func main() {
	livro := Livro{
		titulo: "ABC",
		autor:  "Xuxa",
		preco:  50.0,
	}
	desconto(&livro)
	fmt.Print(livro)
}

func desconto(l *Livro) {
	l.preco = l.preco * 0.90
}
