package main

import "fmt"

func main() {
	var x int
	y := 5

	somaNaturais(&x, y)

	fmt.Println("A soma é igual a: ", x)
}

func somaNaturais(x *int, y int) {
	soma := 0
	for i := 1; i <= y; i++ {
		soma += i
	}
	*x = soma
}
