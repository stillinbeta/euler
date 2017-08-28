package main

import "fmt"

const limit = 1000

func main() {
	sum := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}

	}
	fmt.Println(sum)
}
