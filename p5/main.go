package main

import "fmt"

func main() {
	for i := 20; true; i += 20 {
		if isDivisible(i) {
			fmt.Println(i)
			return
		}

	}

}

func isDivisible(n int) bool {
	for i := 20; i > 1; i-- {
		if n%i != 0 {
			return false
		}
	}
	return true

}
