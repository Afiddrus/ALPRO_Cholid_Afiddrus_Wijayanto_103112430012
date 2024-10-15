package main

// Untuk Input/Output
import "fmt"

func main() {
	// Menampilkan pesan untuk menunjukkan bahwa program akan mencetak bilangan genap dari 1 sampai 50
	fmt.Println("\nBilangan genap dari 1 sampai 50:")

	// Loop dari 1 hingga 50 untuk memeriksa setiap angka
	for i := 1; i <= 50; i++ {
		// Mengecek apakah angka tersebut adalah bilangan genap
		if i%2 == 0 {
			// Jika genap, angka dicetak
			fmt.Print(i, "\t")
		}
	}
}
