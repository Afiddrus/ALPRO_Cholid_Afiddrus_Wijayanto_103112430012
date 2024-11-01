package main

// Import package fmt untuk fungsi input output
import "fmt"

// Program untuk menentukan apakah suatu bilangan adalah faktor dari bilangan yang lain
func main() {
	var x, y int // Deklarasi variabel x dan y sebagai integer (bilangan bulat)

	// Memasukkan nilai untuk x dan y
	fmt.Print("Masukkan bilangan pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan kedua (y): ")
	fmt.Scan(&y)

	// Mengecek apakah x adalah faktor dari y
	if y%x == 0 {
		fmt.Println(true)
	}
	if y%x != 0 {
		fmt.Println(false)
	}

	// Mengecek apakah y adalah faktor dari x
	if x%y == 0 {
		fmt.Println(true)
	}
	if x%y != 0 {
		fmt.Println(false)
	}
}
