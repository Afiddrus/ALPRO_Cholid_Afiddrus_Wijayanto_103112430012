package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan x (x >= y): ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan bilangan y (y > 0): ")
	fmt.Scanln(&y)

	if x < y || y <= 0 {
		fmt.Println("Input tidak valid! bilangan harus sesuai aturan yaitu x >= y dan y > 0.")
		return
	}

	result := 0
	reminder := x

	for reminder >= y {
		reminder -= y
		result++
	}

	fmt.Printf("Hasil dari %d div %d adalah %d\n", x, y, result)
}
