package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik dalam sistem koordinat kartesius
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	var ax, ay, bx, by, cx, cy float64

	// Input koordinat titik A, B, dan C
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scan(&cx, &cy)

	// Menghitung panjang sisi-sisi segitiga
	ab := hitungJarak(ax, ay, bx, by)
	bc := hitungJarak(bx, by, cx, cy)
	ca := hitungJarak(cx, cy, ax, ay)

	// Menentukan sisi terpanjang
	terpanjang := math.Max(ab, math.Max(bc, ca))

	// Output panjang sisi terpanjang dengan dua angka di belakang koma
	fmt.Printf("Panjang sisi terpanjang: %.2f\n", terpanjang)
}
