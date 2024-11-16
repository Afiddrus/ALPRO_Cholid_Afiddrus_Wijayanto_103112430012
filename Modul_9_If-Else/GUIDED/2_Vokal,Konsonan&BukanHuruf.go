package main

import (
	"fmt"
	"strings"
)

func main() {
	//Program untuk menentukan alfabet jika dimasukkan huruf vokal maka akan menampilkan output "vokal", jika konsonan maka output "konsonan", jika bukan huruf maka output "bukan huruf" menggunakan if else
	var huruf string

	fmt.Print("Masukkan Huruf: ")
	fmt.Scanln(&huruf)
	huruf = strings.ToLower(huruf)

	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Println("vokal")
	} else if huruf >= "a" && huruf <= "z" {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}

}
