package main

import (
	"fmt"
)

// Definisikan tipe bentukan untuk menyimpan data investasi saham
type Investasi struct {
	hargaBeli        float64
	hargaJual        float64
	jumlahSaham      int
	totalInvestasi   float64
	totalPenjualan   float64
	keuntunganKotor  float64
	biayaTransaksi   float64
	pajakKeuntungan  float64
	keuntunganBersih float64
}

func main() {
	var inv Investasi

	// Input data investasi saham
	fmt.Println("Informasi Investasi Saham")
	fmt.Print("Harga Beli: Rp ")
	fmt.Scan(&inv.hargaBeli)
	fmt.Print("Harga Jual: Rp ")
	fmt.Scan(&inv.hargaJual)
	fmt.Print("Jumlah Saham: ")
	fmt.Scan(&inv.jumlahSaham)

	// Menghitung nilai-nilai terkait
	inv.totalInvestasi = inv.hargaBeli * float64(inv.jumlahSaham)
	inv.totalPenjualan = inv.hargaJual * float64(inv.jumlahSaham)
	inv.keuntunganKotor = inv.totalPenjualan - inv.totalInvestasi
	inv.biayaTransaksi = 0.002 * inv.totalPenjualan
	inv.pajakKeuntungan = 0.1 * inv.keuntunganKotor
	inv.keuntunganBersih = inv.keuntunganKotor - inv.biayaTransaksi - inv.pajakKeuntungan

	// Menampilkan hasil
	fmt.Printf("\nTotal Investasi Awal: Rp %.2f\n", inv.totalInvestasi)
	fmt.Printf("Total Penjualan: Rp %.2f\n", inv.totalPenjualan)
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", inv.keuntunganKotor)
	fmt.Printf("Biaya Transaksi: Rp %.2f\n", inv.biayaTransaksi)
	fmt.Printf("Pajak Keuntungan: Rp %.2f\n", inv.pajakKeuntungan)
	fmt.Printf("Keuntungan Bersih: Rp %.2f\n", inv.keuntunganBersih)
}
