package main

import "fmt"

func main() {
	var nama_tanaman string
	fmt.Scan(&nama_tanaman)
	switch nama_tanaman {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "calatrava", "bromus":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "cucumbers", "papaya":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "lentils", "broccoli":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "spinach", "kale":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "tomato", "potato":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "zucchini", "eggplant":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	case "pepper", "onion":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	default:
		fmt.Println("Tidak termasuk Tanaman Karnivora")
	}
}
