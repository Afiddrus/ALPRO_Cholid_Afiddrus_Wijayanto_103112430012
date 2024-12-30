package main

import "fmt"

func main() {
	var totalDetik int
	fmt.Println("Masukkan waktu dalam detik:")
	fmt.Scanln(&totalDetik)

	jam := totalDetik / 3600          // 1 jam = 3600 detik
	menit := (totalDetik % 3600) / 60 // Sisa detik setelah dikurangi jam dibagi 60
	detik := totalDetik % 60          // Sisa detik setelah dikurangi menit

	// Menampilkan hasil konversi
	fmt.Printf("%d jam %d menit %d detik\n", jam, menit, detik)
}
