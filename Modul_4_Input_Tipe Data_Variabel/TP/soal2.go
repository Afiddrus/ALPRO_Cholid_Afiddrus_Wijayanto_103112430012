package main

// Import modul fmt yang berisi fungsi-fungsi input-output
import "fmt"

// Fungsi utama
func main() {
	// Deklarasi variabel jam_kerja, upah, jam_normal, jam_lembur dan gaji sebagai float64 (bilangan desimal)
	var jam_kerja, upah, jam_normal, jam_lembur, gaji_bulanan float64

	// Input Jam Kerja
	fmt.Print("Masukkan jam kerja dalam seminggu: ")

	// Mengambil input dan menyimpannya ke variabel jam_kerja
	fmt.Scanln(&jam_kerja)

	// Input Upah
	fmt.Print("Masukkan upah per jam: ")

	// Mengambil input dan menyimpannya ke variabel upah
	fmt.Scanln(&upah)

	// Menghitung jam normal dan jam lembur dengan pernyataan kondisi (if else)
	if jam_kerja > 40 {
		jam_normal = 40
		jam_lembur = jam_kerja - 40
	} else {
		jam_normal = jam_kerja
		jam_lembur = 0
	}

	// Menghitung gaji bulanan (4 minggu)
	gaji_bulanan = (jam_normal*upah + jam_lembur*1.5*upah) * 4

	// Output Gaji per bulan (dengan format float 2 angka dibelakang koma)
	fmt.Printf("Total Gaji Bulanan: %.2f\n", gaji_bulanan)

}
