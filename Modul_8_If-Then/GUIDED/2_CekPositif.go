package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka = ")
	fmt.Scan(&angka)
	if angka > 0 {
		fmt.Println("Positif")
		return
	}
	fmt.Println("Bukan Positif")
}
