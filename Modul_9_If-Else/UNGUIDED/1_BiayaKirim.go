package main

import "fmt"

// PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel
// Program ini akan menghitung biaya jasa dan biaya kirim sesuai dengan ketentuan yang sudah ditetapkan

func main() {
	var beratParsel int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	// Hitung total berat dalam kg dan sisa berat (gram)
	totalBeratKg := beratParsel / 1000 // total berat dalam kg
	sisaBeratGr := beratParsel % 1000  // sisa berat dalam gram

	// Biaya jasa pengiriman: Rp 10.000 per kg
	biayaJasa := float64(totalBeratKg) * 10000

	// Biaya kirim tambahan
	var biayaKirim float64
	if totalBeratKg > 10 {
		// Jika berat lebih dari 10kg, biaya kirim sisa berat digratiskan
		biayaKirim = 0
	} else if sisaBeratGr >= 500 {
		// Jika sisa berat >= 500 gram, biaya tambahan Rp 5 per gram
		biayaKirim = float64(sisaBeratGr) * 5
	} else {
		// Jika sisa berat < 500 gram, biaya tambahan Rp 15 per gram
		biayaKirim = float64(sisaBeratGr) * 15
	}

	// Hitung total biaya
	totalBiaya := biayaJasa + biayaKirim

	// Menampilkan detail output
	fmt.Printf("Detail berat: %d kg + %d gr\n", totalBeratKg, sisaBeratGr)
	fmt.Printf("Detail biaya: Rp. %.0f + Rp. %.0f\n", biayaJasa, biayaKirim)
	fmt.Printf("Total biaya: Rp. %.0f\n", totalBiaya)
}
