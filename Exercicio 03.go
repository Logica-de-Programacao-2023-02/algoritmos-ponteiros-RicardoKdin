package main

import "fmt"

func main() {
	x := "Ricardo Moraes"

	reverso(&x)
	fmt.Print("Reverso: ", x)

}

func reverso(x *string) {
	original := *x
	runes := []rune(original)
	n := len(runes)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*x = string(runes)
}
