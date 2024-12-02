package main

import "fmt"

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&number)

	if number <= 0 {
		fmt.Println("Bilangan harus positif!")
		return
	}

	fmt.Println("Hasil digit dari bilangan (dari kanan ke kiri): ")

	for number > 0 {
		digit := number % 10
		fmt.Println(digit)
		number = number / 10
	}
}
