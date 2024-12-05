package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Masukan harus bilangan bulat positif.")
		return
	}

	count := 0
	selesai := false

	for selesai == false {
		n = n / 10
		count++
		selesai = (n == 0)
	}

	fmt.Println("Banyak digit:", count)
}
