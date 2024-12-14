package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor-faktor dari ", n, " adalah: ")

	//Perulangan untuk mencari sebuah faktor dari bilangan yang diinputkan
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}

}
