package main

import "fmt" // Import package fmt untuk input dan output

func main() {
	// Deklarasi variabel x, y, dan sum sebagai integer
	var x, y, sum int

	// Input nilai X dan Y
	fmt.Print("Masukkan Nilai X: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan Nilai Y: ")
	fmt.Scanln(&y)

	// Inisialisasi nilai sum ke 0
	sum = 0

	// Menjumlahkan bilangan dari x sampai dengan y
	for i := x; i <= y; i++ {
		sum += i
	}

	// Print atau tampilkan hasil penjumlahan
	fmt.Println("Hasil Penjumlahan:", sum)
}
