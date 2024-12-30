package main

import "fmt"

func main() {
	var char byte
	fmt.Print("Masukkan satu karakter huruf: ")
	fmt.Scanf("%c", &char)

	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			fmt.Println("Bukan konsonan")
		} else {
			fmt.Println("Konsonan")
		}
	} else {
		fmt.Println("Bukan konsonan")
	}
}
