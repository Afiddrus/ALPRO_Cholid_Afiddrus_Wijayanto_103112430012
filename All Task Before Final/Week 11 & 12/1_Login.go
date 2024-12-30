package main

import "fmt"

func main() {
	usernameBenar := "admin"
	passwordBenar := "admin"

	gagalLogin := 0

	var username, password string

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username == usernameBenar && password == passwordBenar {
			break
		} else {
			gagalLogin++
			fmt.Println("Username atau password salah. Silakan coba lagi.")
		}
	}

	fmt.Printf("Jumlah kegagalan login: %d\n", gagalLogin)
	fmt.Println("Login berhasil")
}
