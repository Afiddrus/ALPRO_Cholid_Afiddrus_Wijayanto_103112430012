package main

import (
	"fmt"
)

func main() {
	var a, b, hasil int

	fmt.Print("Masukkan Bilangan Pertama: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan Bilangan Kedua: ")
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		hasil = hasil + a
	}
	print("Hasil: ", hasil)
}
