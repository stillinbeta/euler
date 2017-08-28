package main

import (
	"fmt"
	"math"
)

const size = 9

func main() {
	for i := size; i > 0; i-- {
		c := make(chan int, 0)
		go generatePanDigit(i, c)
		for n := range c {
			if isPrime(n) {
				fmt.Printf("%d\n", n)
				return
			}
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

func generatePanDigit(n int, c chan int) {
	arr := makeRange(n)
	generatePanDigitFromArr([]int{}, arr, c)
	close(c)
}

func generatePanDigitFromArr(locked, free []int, c chan int) {
	if len(free) == 0 {
		num := 0
		for i := 0; i < len(locked); i++ {
			num += locked[i] * int(math.Pow(10, float64(len(locked)-i)-1))
		}
		c <- num
	}

	for i := 0; i < len(free); i++ {
		freep := []int{}
		// haskell in go >_<
		freep = append(freep, free[:i]...)
		freep = append(freep, free[i+1:]...)
		generatePanDigitFromArr(append(locked, free[i]), freep, c)
	}
}

func makeRange(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i
	}
	return arr
}
