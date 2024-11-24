package main

import "fmt"

func main() {
	const password = "rahasia123"
	var percobaan = 3
	var isLogined = false

	fmt.Println("Selamat datang di sistem login")

	for percobaan > 0 {
		var inputPassword string
		fmt.Printf("Masukkan password(%d kesempatan tersisa): ", percobaan)
		fmt.Scanln(&inputPassword)

		if inputPassword == password {
			isLogined = true
			fmt.Println("Login berhasil")
			break
		} else {
			percobaan--
			if percobaan > 0 {
				fmt.Println("Password salah. Silakan coba lagi")
			}
		}
	}
	if !isLogined {
		fmt.Println("Login ditolak")
	}
}
