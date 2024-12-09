package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&num)

	fmt.Printf("Bilangan prima dari 1 sampai %d adalah:\n", num)

	for i := 2; i <= num; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
