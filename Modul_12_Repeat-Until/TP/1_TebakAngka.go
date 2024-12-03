package main

import "fmt"

func main() {
	const secretNumber = 7 // Angka rahasia
	var guess int

	fmt.Println("Tebak angka rahasia antara 1 hingga 10!")

	for {
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scanln(&guess)

		// Cek tebakan benar atau tidak
		if guess == secretNumber {
			fmt.Println("Selamat, tebakan Anda benar!")
			break // Keluar dari loop karena tebakan benar
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}
