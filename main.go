package main

import (
	"fmt"
)

func main() {
	var alas, tinggi, luas float64

	// Input alas segitiga
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scan(&alas)

	// Input tinggi segitiga
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	// Hitung luas segitiga
	luas = 0.5 * alas * tinggi

	// Tampilkan hasil
	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luas)
}
