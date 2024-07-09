package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1073741824) // 2^30
	b := big.NewInt(33554432)   // 2^25
	c := new(big.Int)

	fmt.Println("a + b = ", c.Add(a, b)) // a + b =  1107296256
	fmt.Println("a - b = ", c.Sub(a, b)) // a - b =  1040187392
	fmt.Println("a / b = ", c.Div(a, b)) // a / b =  32
	fmt.Println("a * b = ", c.Mul(a, b)) // a * b =  36028797018963968
}
