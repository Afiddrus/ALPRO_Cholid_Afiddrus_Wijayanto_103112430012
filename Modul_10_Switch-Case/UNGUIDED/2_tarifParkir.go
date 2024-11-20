package main

import (
	"fmt"
	"strings"
)

func main() {
	var kendaraan string
	var durasi float64

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scanln(&kendaraan)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scanln(&durasi)

	kendaraan = strings.ToLower(kendaraan)

	if durasi < 1 {
		durasi = 1
	}

	var tarifPerJam int
	switch kendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid. Pilih motor, mobil, atau truk.")
		return
	}

	totalBiaya := tarifPerJam * int(durasi)

	fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}
