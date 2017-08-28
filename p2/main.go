package main

import "fmt"

const max = 4000000

func main() {

	a, b := 1, 1
	sum := 0
	for a < max {
		if a%2 == 0 {
			sum += a
		}
		a, b = b, a+b
	}
	fmt.Println(sum)

}
