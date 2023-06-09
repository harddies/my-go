package main

import "fmt"

func main() {
	fmt.Println(isPrimeNumber(23))
}

func isPrimeNumber(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
