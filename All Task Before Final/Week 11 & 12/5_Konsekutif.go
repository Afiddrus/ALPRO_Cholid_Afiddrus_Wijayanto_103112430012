package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&bilangan)

	konsekutif := true

	var digitSebelumnya int = -1

	for bilangan > 0 {
		digitSekarang := bilangan % 10

		if digitSebelumnya != -1 {
			if nilaiMutlak(digitSekarang-digitSebelumnya) != 1 {
				konsekutif = false
				break
			}
		}

		digitSebelumnya = digitSekarang
		bilangan /= 10
	}

	fmt.Println("Apakah bilangan konsekutif?", konsekutif)
}

func nilaiMutlak(angka int) int {
	if angka < 0 {
		return -angka
	}
	return angka
}
