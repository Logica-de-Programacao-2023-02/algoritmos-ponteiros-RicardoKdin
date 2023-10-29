package main

import "fmt"

func main() {
	x := 7

	parImpar(&x)
	if x == 0 {
		fmt.Print("O número é par")
	} else {
		fmt.Println("O número é ímpar")
	}
}

func parImpar(x *int) {
	if *x%2 == 0 {
		*x = 0
	} else {
		*x = 1
	}
}
