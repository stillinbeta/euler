package main

import (
	"fmt"
	"math"
)

func main() {
	max := 0
	// I think there's a more efficient solution
	// but I can't find it
	for i := 999; i > 0; i-- {
		for j := 1; j <= 999; j++ {
			prod := i * j
			if prod > max && isPalindrome(prod) {
				max = prod
			}

		}

	}
	fmt.Println(max)

}

func isPalindrome(n int) bool {
	size := int(math.Log10(float64(n)) + 1)
	// fmt.Println(size)
	for i := 0; i < size/2; i++ {
		powRight := int(math.Pow(10, float64(i)) + 0.5)
		powLeft := int(math.Pow(10, float64(size-i-1)) + 0.5)
		// fmt.Printf("l: %d, r: %d\n", powLeft, powRight)

		if (n / powLeft % 10) != (n / powRight % 10) {
			return false
		}

	}
	return true
}
