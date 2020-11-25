package main

import "fmt"

func main() {
	primes := []int{}
	primes = append(primes, 1)
	primes = append(primes, 2)
	primes = append(primes, 5)
	fmt.Println(primes)
}