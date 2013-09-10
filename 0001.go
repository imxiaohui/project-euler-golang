// If we list all the natural numbers below 10 that are multiples of 3 or
// 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

// Easy solution; not used
func p1easy(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func p1fast(target int) int {
	// sdb: SumDivisibleBy
	sdb := func(n int) int {
		p := target / n
		return n * (p * (p + 1)) / 2
	}
	return sdb(3) + sdb(5) - sdb(15)
}
