package main

import "fmt"

func main() {
    var n int
    
    fmt.Print("Masukkan sebuah bilangan positif: ")
    fmt.Scan(&n)
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if j == i || j == (n-i+1) {
                fmt.Printf("%d ", i)
            } else {
                fmt.Print("  ")
            }
        }
        fmt.Println()
    }
}