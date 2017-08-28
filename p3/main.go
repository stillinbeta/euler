package main

import "fmt"

const input = 600851475143

func main() {
	for i := input / 2; i > 0; i-- {
		if input%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
		if i%100000000 == 0 {
			fmt.Println(i)

		}
	}
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true

}
