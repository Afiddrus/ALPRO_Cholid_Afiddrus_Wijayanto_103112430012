package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Masukkan angka A: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka B: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		a++
		fmt.Print(i, " ")
	}
}
