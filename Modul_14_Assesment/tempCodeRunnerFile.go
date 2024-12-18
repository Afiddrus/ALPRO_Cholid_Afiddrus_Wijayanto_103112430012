package main

import "fmt" // Import package fmt untuk fungsi input dan output

// Fungsi utama
func main() {
	var a_103112430012, b, c int // Deklarasi variabel a, b, dan c yang digunakan untuk input oleh user yang bertipe data int (bilangan bulat)

	// Meminta input untuk variabel a
	fmt.Print("Masukkan panjang sisi a: ")
	fmt.Scan(&a_103112430012)

	// Meminta input untuk variabel b
	fmt.Print("Masukkan panjang sisi b: ")
	fmt.Scan(&b)

	// Meminta input untuk variabel c
	fmt.Print("Masukkan panjang sisi c: ")
	fmt.Scan(&c)

	// Perulangan untuk mengecek apakah segitiga sama sisi, sama kaki, sembarang, atau bukan segitiga
	if a_103112430012 == b && b == c { // Jika a dan b dan c memiliki nilai yang sama maka merupakan segitiga sama sisi
		fmt.Println("Segitiga sama sisi")
	} else if a_103112430012 == b || b == c || a_103112430012 == c { // Jika a atau b atau c memiliki nilai yang sama maka merupakan segitiga sama kaki
		fmt.Println("Segitiga sama kaki")
	} else if c-b == 1 || b-a_103112430012 == 1 || c-a_103112430012 == 1 { // Jika c - b atau b - a bernilai 1 maka merupakan segitiga sembarang
		fmt.Println("Segitiga sembarang")
	} else if a_103112430012%2 != 0 && c%2 != 0 { // jika a merupakan bilangan ganjil dan c merupakan bilangan ganjil maka merupakan segitiga siku-siku
		fmt.Println("Segitiga siku-siku")
	} else {
		fmt.Println("Bukan segitiga")
	}

}
