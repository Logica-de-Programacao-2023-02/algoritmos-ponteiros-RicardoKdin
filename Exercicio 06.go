package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func main() {
	l1 := Livro{
		Titulo: "ABC",
		Autor:  "FULANO",
	}
	l2 := Livro{
		Titulo: "XXX",
		Autor:  "Anônimo",
	}
	changeBook(&l1)
	changeBook(&l2)

	fmt.Println(l1)
	fmt.Println(l2)
}

func changeBook(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Titulo = "Desconhecido"
	}
}
