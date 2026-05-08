package main

import "fmt"

func main() {
	// 1. Hello World Dasar
	fmt.Println("Halo, dunia! Memulai eksekusi kode Golang pertama.")

	// 2. Deklarasi Variabel Eksplisit
	var angkaPertama int = 15
	var angkaKedua int = 5

	// 3. Deklarasi Variabel Shorthand (Paling sering digunakan)
	operasi := "Penjumlahan & Pengurangan"
	hasilJumlah := angkaPertama + angkaKedua
	hasilKurang := angkaPertama - angkaKedua

	// 4. "Hello World" Steroid: Menampilkan variabel dengan Printf (Print Format)
	// %s untuk String, %d untuk Angka (Digit/Integer), \n untuk baris baru (enter)
	fmt.Printf("\n--- Hasil Operasi %s ---\n", operasi)
	fmt.Printf("Angka yang digunakan: %d dan %d\n", angkaPertama, angkaKedua)
	fmt.Printf("Hasil Penjumlahan: %d\n", hasilJumlah)
	fmt.Printf("Hasil Pengurangan: %d\n", hasilKurang)
}
