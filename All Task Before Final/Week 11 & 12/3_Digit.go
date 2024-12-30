package main

import "fmt"

func main() {
	var X int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&X)

	total := 0

	first := true

	for X > 0 {
		digit := X % 10

		if !first {
			fmt.Print(" ")
		}
		fmt.Print(digit)

		total += digit

		X /= 10

		first = false
	}

	fmt.Println()
	fmt.Println(total)
}
