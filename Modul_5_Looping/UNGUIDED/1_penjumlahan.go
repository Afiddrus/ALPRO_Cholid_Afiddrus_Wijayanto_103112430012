package main

// Mengimpor paket fmt untuk fungsi format input/output.
import "fmt"

// Fungsi utama, tempat eksekusi program dimulai.
func main() {
	var input int // Mendeklarasikan variabel 'input' bertipe integer.

	fmt.Print("Masukkan angka: ") // Mencetak pesan untuk meminta input dari pengguna.
	fmt.Scan(&input)              // Membaca input dari pengguna dan menyimpannya dalam variabel 'input'.

	// Memeriksa apakah input kurang dari atau sama dengan 0
	if input <= 0 {
		fmt.Println("Angka harus lebih besar dari 0") // Menampilkan pesan jika input tidak valid.
		return                                        // Menghentikan eksekusi program.
	}

	tambah := 0 // Mendeklarasikan variabel 'tambah' untuk menyimpan hasil penjumlahan, diinisialisasi dengan 0.

	// Loop untuk menghitung penjumlahan dari 1 sampai dengan input.
	for i := 0; i <= input; i++ {
		tambah += i // Menambahkan nilai i ke variabel 'tambah' pada setiap iterasi.
	}
	// Mencetak hasil penjumlahan.
	print("Hasil Penjumlahan dari 1 sampai ke ", input, " adalah:", tambah)
}
