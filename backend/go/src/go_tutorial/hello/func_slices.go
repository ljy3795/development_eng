package main

import "fmt"

// check the number is prime
func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			// condition is true then break
			return false
		}
	}
	return true
}

// returning a slice
func getPrimes(n int, a []int) []int {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			a = append(a, i) //slice append
		}
	}
	return a
}

func main() {
	num_v := 10

	fmt.Println(isPrime(num_v))

	max_v := 30
	ans_v := []int{}
	fmt.Println(getPrimes(max_v, ans_v))
	fmt.Println(getPrimes(max_v, ans_v)[2])

}
