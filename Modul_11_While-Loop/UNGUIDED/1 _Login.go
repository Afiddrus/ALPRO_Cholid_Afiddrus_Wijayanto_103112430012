package main

import "fmt"

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"

	var username, password string
	var failedAttempts int

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)

		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username == correctUsername && password == correctPassword {
			fmt.Printf("%d percobaan gagal login\n", failedAttempts)
			return
		} else {
			fmt.Println("Username atau password salah, coba lagi.")
			failedAttempts++
		}
	}
}
