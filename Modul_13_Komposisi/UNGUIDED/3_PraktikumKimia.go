package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Percobaan struct {
	gelas []string
}

type HasilPraktikum struct {
	percobaan []Percobaan
}

func NewHasilPraktikum() *HasilPraktikum {
	return &HasilPraktikum{
		percobaan: make([]Percobaan, 0),
	}
}

func (h *HasilPraktikum) TambahPercobaan(gelas []string) {
	h.percobaan = append(h.percobaan, Percobaan{gelas: gelas})
}

func (h *HasilPraktikum) CekHasil() bool {
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}

	var i int
	for i = 0; i < len(h.percobaan); i++ {
		if len(h.percobaan[i].gelas) != 4 {
			return false
		}

		var j int
		for j = 0; j < 4; j++ {
			if h.percobaan[i].gelas[j] != urutanBenar[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	hasil := NewHasilPraktikum()
	scanner := bufio.NewScanner(os.Stdin)

	var i int
	for i = 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		scanner.Scan()
		input := scanner.Text()

		warna := strings.Fields(input)

		hasil.TambahPercobaan(warna)
	}

	fmt.Printf("BERHASIL: %t\n", hasil.CekHasil())
}
