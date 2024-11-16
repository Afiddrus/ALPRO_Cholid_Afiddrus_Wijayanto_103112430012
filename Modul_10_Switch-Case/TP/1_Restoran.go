package main

// Import package fmt untuk fungsi input dan output
import "fmt"

// Fungsi utama
func main() {
	// Menampilkan menu
	fmt.Println("Menu Restoran Cepat Saji")
	fmt.Println("1. Burger - Rp25,000")
	fmt.Println("2. Fried Chicken - Rp20,000")
	fmt.Println("3. French Fries - Rp15,000")
	fmt.Println("4. Soft Drink - Rp10,000")
	fmt.Println("5. Coffee - Rp15,000")

	var input string // Dekalarasi variabel input dengan tipe data string
	// Input kode item
	fmt.Print("Masukkan kode item (1-5): ")
	fmt.Scanln(&input)

	//Fungsi switch-case
	switch input {
	case "1": // Jika user menginputkan 1 maka menampilkan menu burger
		fmt.Println("Menu yang dipilih: Burger - Rp25,000")

	case "2": // Jika user menginputkan 2 maka menampilkan menu fried chicken
		fmt.Println("Menu yang dipilih: Fried Chicken - Rp20,000")

	case "3": // Jika user menginputkan 3 maka menampilkan menu french fries
		fmt.Println("Menu yang dipilih: French Fries - Rp15,000")

	case "4": // Jika user menginputkan 4 maka menampilkan menu soft drink
		fmt.Println("Menu yang dipilih: Soft Drink - Rp10,000")

	case "5": // Jika user menginputkan 5 maka menampilkan menu coffee
		fmt.Println("Menu yang dipilih: Coffee - Rp15,000")

	default: // Selain itu maka akan menampilkan menu tidak valid
		fmt.Println("Menu yang dipilih: Kode menu tidak valid")
	}
}
