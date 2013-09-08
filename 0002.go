// Each new term in the Fibonacci sequence is generated by adding the
// previous two terms. By starting with 1 and 2, the first 10 terms will
// be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not
// exceed four million, find the sum of the even-valued terms.

package main

import "fmt"

func p2(n int) int {
	sum, a, b, c := 0, 1, 1, 2 // c=a+b
	for c < n {
		sum += c
		a = b + c
		b = c + a
		c = a + b
	}
	return sum
}

func main() {
	fmt.Println(p2(4000000))
}
