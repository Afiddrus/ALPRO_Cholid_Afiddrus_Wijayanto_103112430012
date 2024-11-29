package main

import "fmt"

func main() {
	var token string
	maxPercobaan := 3
	percobaan := 0

	for percobaan < maxPercobaan {
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&token)

		if token == "123456abcde" {
			fmt.Println("Selamat, Anda berhasil login!")
			return
		}

		percobaan++
		if percobaan < maxPercobaan {
			fmt.Println("Maaf, Password salah")
		}
	}
	fmt.Println("Anda tidak bisa login, kesempatan anda sudah 3x login habis!")

}
