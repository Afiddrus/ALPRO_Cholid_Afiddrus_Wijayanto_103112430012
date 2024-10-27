package main

// import format untuk input dan output
import "fmt"

// deklarasi fungsi
func main() {
	// deklarasi variabel jumlah semester dan skor eprt dengan tipe data integer
	var JumlahSemester, skorEPrT int

	// Input jumlah semester
	fmt.Print("Masukkan Jumlah Semester: ")
	fmt.Scan(&JumlahSemester)

	// Input skor EPrT
	fmt.Print("Masukkan Skor EPrT: ")
	fmt.Scan(&skorEPrT)

	// deklarasi variabel untuk menentukan apakah mahasiswa lulus cumlaude dengan tipe data boolean untuk menentukan nilai true atau false
	var cumlaude bool

	// perhitungan apakah mahasiswa lulus cumlaude
	cumlaude = JumlahSemester <= 8 && skorEPrT >= 500

	// mengecek apakah mahasiswa memenuhi syarat untuk lulus cumlaude
	if cumlaude == true {
		// Jika mahasiswa lulus cumlaude, cetak pesan kelulusan
		fmt.Print("mahasiswa lulus cumlaude dengan kuliah selama ", JumlahSemester, " semester dan EPrT ", skorEPrT)
	} else if JumlahSemester > 8 && skorEPrT < 500 {
		// Jika jumlah semester lebih dari 8 dan skor EPrT kurang dari 500
		fmt.Print("mahasiswa tidak cumlaude karena kuliah hingga ", JumlahSemester, " semester dan skor EPrT = ", skorEPrT)
	} else if JumlahSemester > 8 {
		// Jika hanya jumlah semester lebih dari 8
		fmt.Print("mahasiswa tidak cumlaude karena kuliah hingga ", JumlahSemester, " semester")
	} else if skorEPrT < 500 {
		// Jika hanya skor EPrT kurang dari 500
		fmt.Print("mahasiswa tidak cumlaude karena skor EPrT = ", skorEPrT)
	}
}
