package main

import "fmt"

func main() {
	var a int
	var b bool

	fmt.Scan(&a)
	b = false
	if a < 0 && a%2 == 0 {
		b = true
	}
	fmt.Println(b)
}
