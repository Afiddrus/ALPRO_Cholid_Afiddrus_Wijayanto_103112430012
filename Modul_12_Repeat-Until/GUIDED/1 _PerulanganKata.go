package main

import "fmt"

func main() {
	var word string
	var repetitions int

	fmt.Scan(&word, &repetitions)

	counter := 0

	for counter < repetitions {
		fmt.Println(word)
		counter++
	}
}
