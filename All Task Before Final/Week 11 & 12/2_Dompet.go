package main

import "fmt"

func main() {
	saldo := 0

	var transaksi int

	for {
		fmt.Print("Masukkan transaksi (0 untuk berhenti): ")
		fmt.Scanln(&transaksi)

		if transaksi == 0 {
			break
		}

		saldo += transaksi
	}

	fmt.Println("Saldo akhir dompet: ", saldo)
}
