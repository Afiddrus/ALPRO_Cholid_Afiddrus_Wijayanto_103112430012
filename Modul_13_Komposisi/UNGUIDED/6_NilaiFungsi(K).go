package main

import (
	"fmt"
	"math"
)

func main() {
	var k int

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	result := 1.0

	for i := 0; i <= k; i++ {
		numerator := math.Pow(float64(4*k+2), 2)
		denumerator := float64((4 * k) * (4*k + 3))
		result *= numerator / denumerator
	}

	fmt.Printf("Hasil dari operasi fungsi = %.10f\n", result)

}
