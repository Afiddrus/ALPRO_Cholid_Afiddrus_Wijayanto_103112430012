package main

// Import modul fmt yang berisi fungsi-fungsi input-output
import (
	"fmt"
)

// Fungsi utama
func main() {
	// Deklarasi variabel pi dengan tipe data konstanta
	const pi = 3.14

	// Deklarasi variabel r(jari-jari) dengan tipe data float64 (bilangan desimal)
	var r float64

	// Input jari-jari
	fmt.Print("Masukkan jari-jari = ")

	// Mengambil input dan menyimpannya ke variabel r(jari-jari)
	fmt.Scanln(&r)

	// Rumus menghitung luas lingkaran
	luas := pi * r * r

	// Rumus menghitung keliling lingkaran
	keliling := 2 * pi * r

	// Output hasil luas dan keliling lingkaran
	fmt.Println("Luas dari lingkaran tersebut adalah:", luas)
	fmt.Println("Keliling dari lingkaran tersebut adalah:", keliling)
}
