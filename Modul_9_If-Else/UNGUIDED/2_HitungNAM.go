// a. Jika nam diberikan adalah 80.1,keluaran dari program tersebut adalah error dikarenakan terdapat kesalahan dalam tipe data. Program tersebut menyimpan string ("A", "AB", dll.) ke dalam variabel nam yang seharusnya bertipe float64. Akibatnya. kondisi perbandingan tidak sesuai dan tidak berfungsi.
// b. Kesalahan dari program tersebut: 1. Tipe data salah: Variabel nam diubah menjadi string untuk menyimpan nilai huruf (A, AB, dll), seharusnya tetap bertipe float64 untuk perhitungan. 2.Logika kondisi tidak tepat: Seharusnya kondisi if-else diperiksa secara berurutan dari nilai tertinggi hingga terendah 3. Penggunaan variabel nmk tidak sesuai: Program menyalahgunakan variabel nam untuk menyimpan nilai huruf, padahal nmk seharusnya yang menyimpan hasil penilaian. Alur yang benar adalah: 1. Menerima input nam (nilai akhir mata kuliah). 2. Berdasarkan rentang nilai nam, tentukan nilai huruf nmk. 3. Cetak nilai nmk sesuai dengan ketentuan.4
// c. Program diperbaiki dengan cara: 1. Menjaga tipe data nam tetap float64 untuk menerima input numerik 2. Menggunakan variabel nmk untuk menyimpan nilai huruf 3. Menggunakan struktur if-else untuk memeriksa rentang nilai dari yang terbesar hingga terkecil. Berikut adalah perbaikannya:

package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah: ", nmk)
}
