package main

import (
	"fmt"
	"math/big"
)

const fact = 100

func main() {
	base := big.NewInt(1)
	mul := big.NewInt(1)
	for i := 1; i <= fact; i++ {
		base = base.Mul(base, mul.SetInt64(int64(i)))
	}

	acc := big.NewInt(0)
	ten := big.NewInt(10)
	pow := big.NewInt(0)
	mod := big.NewInt(0)

	// No logarithim function in math/big
	size := len(base.String())

	for i := size; i >= 0; i-- {
		pow := pow.Exp(ten, mul.SetInt64(int64(i)), nil)
		base.DivMod(base, pow, mod)
		acc.Add(acc, base)
		base.Set(mod)
	}
	fmt.Println(acc.String())

}
