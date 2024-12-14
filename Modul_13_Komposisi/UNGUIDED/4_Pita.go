package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PitaBunga struct {
	isi    string
	jumlah int
}

func NewPitaBunga() *PitaBunga {
	return &PitaBunga{
		isi:    "",
		jumlah: 0,
	}
}

func (p *PitaBunga) TambahBunga(bunga string) {
	if p.isi == "" {
		p.isi = bunga
	} else {
		p.isi += " - " + bunga
	}
	p.jumlah++
}

func InputDenganN() {
	var n int
	fmt.Print("N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Pita: ")
		return
	}

	pita := NewPitaBunga()
	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d: ", i)
		scanner.Scan()
		bunga := scanner.Text()
		bunga = strings.TrimSpace(bunga)

		if bunga != "" {
			pita.TambahBunga(bunga)
		}
	}

	fmt.Printf("Pita: %s\n", pita.isi)
}

func InputSampaiSelesai() {
	pita := NewPitaBunga()
	scanner := bufio.NewScanner(os.Stdin)
	i := 1

	for {
		fmt.Printf("Bunga %d: ", i)
		scanner.Scan()
		bunga := scanner.Text()
		bunga = strings.TrimSpace(bunga)

		if bunga == "SELESAI" {
			break
		}

		if bunga != "" {
			pita.TambahBunga(bunga)
			i++
		}
	}

	fmt.Printf("Pita: %s\n", pita.isi)
	fmt.Printf("Bunga: %d\n", pita.jumlah)
}

func main() {
	// InputDenganN()     // Untuk versi dengan input N
	InputSampaiSelesai() // Untuk versi dengan input sampai SELESAI
}
