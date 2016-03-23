//Q6. (0) Average
//1. Write a function that calculates the average of a float64 slice.
package main

import "fmt"

func switcha(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	a, b := switcha(5,4)
	fmt.Println(a, b)
}


