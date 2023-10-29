package main

import "fmt"

func main() {
	x := 20
	err := change(&x)
	fmt.Println(err)
}

func change(x *int) error {
	if x != nil {
		*x = 10
		return nil
	} else {
		return fmt.Errorf("invalid pointer")
	}
}
