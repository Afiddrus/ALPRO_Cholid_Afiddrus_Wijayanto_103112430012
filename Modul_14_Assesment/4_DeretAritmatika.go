package main

import "fmt" // Package fmt untuk input dan output

// Fungsi utama
func main() {
	var n_103112430012 int // Deklarasi variabel a bertipe int (bilangan bulat) yang menjadi jumlah elemen dalam deret
	var a []int            // Array untuk menyimpan

	// Meminta input untuk variabel n dari user
	fmt.Print("Masukkan jumlah elemen deret: ")
	fmt.Scan(&n_103112430012)

	// Jika n < 3 maka deret tersebut tidak valid, jika n >= 3 maka akan melakukan perhitungan dan menampilkan output
	if n_103112430012 < 3 {

		fmt.Println("Deret tidak valid")
	} else {
		fmt.Print("Masukkan elemen-elemen deret: ")
		for i := 0; i < n_103112430012; i++ {
			var elemen int
			fmt.Scan(&elemen)
			a = append(a, elemen)
		}

		// Melakukan pengecekan deret aritmatika jika ada suku yang tidak sesuai maka akan menampilkan deret tidak valid
		var selisih = a[1] - a[0]
		for i := 2; i < n_103112430012; i++ {
			if a[i]-a[i-1] != selisih {
				fmt.Println("Deret tidak valid")
				return
			}
		}

		var jumlah int
		for i := 0; i < n_103112430012; i++ {
			jumlah += a[i]
		}

		fmt.Printf("Jumlah total: %d\n", jumlah)

	}
}
