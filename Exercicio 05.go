package main

import (
	"fmt"
	"math"
)

func main() {
	n := 3.0

	media(&n)
	fmt.Print("Média: ", n)

}

func media(x *float64) {
	*x = (*x + math.Pi) / 2
}
