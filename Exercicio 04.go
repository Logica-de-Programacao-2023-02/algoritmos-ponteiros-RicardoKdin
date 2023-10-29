package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 1234
	updateDigits(&x)
	fmt.Println(x)
}

func updateDigits(x *int) {
	xStr := strconv.Itoa(*x)

	last := xStr[len(xStr)-1]
	secondLast := xStr[len(xStr)-2]

	lastI, _ := strconv.Atoi(string(last))
	secondLastI, _ := strconv.Atoi(string(secondLast))

	*x = lastI + secondLastI
}
