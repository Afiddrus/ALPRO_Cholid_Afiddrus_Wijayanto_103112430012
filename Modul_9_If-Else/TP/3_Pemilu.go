package main

// Import package fmt untuk fungsi input output dan package strings untuk memanipulasi string
import (
	"fmt"
	"strings"
)

// Fungsi utama
func main() {
	var umur int               // Deklarasi variabel umur dengan tipe data integer
	var kewarganegaraan string // Deklarasi variabel kewarganegaraan dengan tipe data string

	// Input umur
	fmt.Print("Masukkan umur Anda: ")
	fmt.Scanln(&umur)

	// Input kewarganegaraan
	fmt.Print("Masukkan kewarganegaraan Anda (contoh: WNI): ")
	fmt.Scanln(&kewarganegaraan)

	// Ubah input kewarganegaraan menjadi huruf besar agar tidak case-sensitive
	kewarganegaraan = strings.ToUpper(kewarganegaraan)

	// Cek syarat pemilu
	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		// Tampilkan alasan yang membuat user tidak memenuhi syarat
		if umur < 17 && kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda di bawah 17 tahun dan Anda bukan WNI")
		} else if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda di bawah 17 tahun")
		} else if kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena Anda bukan WNI")
		}
	}
}
