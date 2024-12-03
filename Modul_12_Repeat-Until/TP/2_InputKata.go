package main

import (
	"fmt"
	"strings" // package untuk mengatasi tidak peka huruf besar/kecil
)

func main() {
	var input string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scanln(&input)

		// Cek apakah input adalah "telkom" (tidak peka huruf besar/kecil)
		if strings.ToLower(input) == "telkom" {
			fmt.Println("Program selesai.")
			break // Keluar dari loop karena input adalah "telkom"
		} else {
			fmt.Printf("Anda mengetik: %s\n", input)
		}
	}
}
