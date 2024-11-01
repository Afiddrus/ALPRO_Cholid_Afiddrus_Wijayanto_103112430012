package main

// Import package fmt untuk fungsi input output
import "fmt"

// Program untuk mengecek apakah sebuah bilangan yang diinputkan merupakan bilangan genap negatif, jika bukan maka menampilkan teks bukan dan jika iya maka menampilkan teks genap negatif
func main() {

	var angka int // deklarasi variabel angka bertipe integer (bilangan bulat)

	fmt.Print("Masukkan angka = ") // menampilkan teks masukkan angka

	fmt.Scan(&angka) // membaca dan menyimpan angka

	// percabangan untuk mengecek apakah angka adalah bilangan genap negatif
	if angka%2 == 0 && angka < 0 {
		fmt.Println("Angka tersebut adalah bilangan genap negatif")
	} else {
		fmt.Println("Bukan bilangan genap negatif")
	}
}
