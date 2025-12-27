package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("92233720368547758075", 10)
	b.SetString("122233720368547758078", 10)

	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", sub)
	fmt.Println("a * b =", mul)
	fmt.Println("a / b =", div)
}
