package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Sistem Integrator Rutinitas ===")

	// 1. Slices: Menyimpan daftar rutinitas (Bisa ditambah dinamis)
	rutinitas := []string{
		"Belajar Golang Sesi 2",
		"Review Algoritma",
		"Membaca Buku",
		"Olahraga Ringan",
	}

	// Menambah data baru ke dalam Slice
	rutinitas = append(rutinitas, "Istirahat & Meditasi")

	// 2. Maps: Menyimpan detail kategori atau tingkat prioritas tugas
	prioritasTugas := map[string]string{
		"Belajar Golang Sesi 2": "Tinggi",
		"Review Algoritma":      "Sedang",
		"Membaca Buku":          "Rendah",
		"Olahraga Ringan":       "Sedang",
		"Istirahat & Meditasi":  "Tinggi",
	}

	// 1. Modifikasi Map: Key (string) memetakan ke Value (int) untuk durasi menit
	durasiTugas := map[string]int{
		"Belajar Golang Sesi 2": 120,
		"Review Algoritma":      90,
		"Membaca Buku":          45,
		"Olahraga Ringan":       30,
		"Istirahat & Meditasi":  15,
	}

	// 3. Control Flow (For Range): Menampilkan semua tugas beserta prioritasnya
	fmt.Println("\nDaftar Rutinitas Hari Ini:")
	for index, tugas := range rutinitas {
		// Mengambil nilai dari Map berdasarkan key (tugas)
		prioritas := prioritasTugas[tugas]
		fmt.Printf("%d. %s (Prioritas: %s)\n", index+1, tugas, prioritas)
	}

	// 4. Control Flow (If/Else & Switch): Logika Eliminasi Decision Fatigue
	fmt.Println("\nMemilihkan rutinitas acak untuk Anda...")

	// Mengatur seed random agar hasilnya berbeda tiap kali dijalankan
	rand.Seed(time.Now().UnixNano())
	taskPriorityIndex := rand.Intn(len(rutinitas))
	tugasTerpilih := rutinitas[taskPriorityIndex]
	tingkatPrioritas := prioritasTugas[tugasTerpilih]

	fmt.Printf("=> Rutinitas Terpilih: %s\n", tugasTerpilih)

	// Penggunaan Switch tanpa break
	switch tingkatPrioritas {
	case "Tinggi":
		fmt.Println("Saran: Kerjakan sekarang juga! Jangan ditunda.")
	case "Sedang":
		fmt.Println("Saran: Boleh dikerjakan setelah tugas prioritas tinggi selesai.")
	case "Rendah":
		fmt.Println("Saran: Kerjakan di waktu luang untuk relaksasi.")
	default:
		fmt.Println("Saran: Status prioritas tidak diketahui.")
	}

	// Memilih tugas secara acak
	rand.Seed(time.Now().UnixNano())
	taskDuration := rand.Intn(len(rutinitas))
	tugasTerpilihDurasi := rutinitas[taskDuration]

	// Mengambil durasi (int) dari Map berdasarkan tugas yang terpilih
	waktu := durasiTugas[tugasTerpilihDurasi]

	fmt.Printf("\n=> Rutinitas Terpilih: %s\n", tugasTerpilihDurasi)
	fmt.Printf("=> Estimasi Durasi: %d menit\n", waktu)

	// Modifikasi Logika If/Else
	if waktu > 60 {
		fmt.Println("Status: Siapkan kopi dan air putih! ☕💧")
	} else if waktu >= 30 && waktu <= 60 {
		fmt.Println("Status: Durasi standar. Tetap fokus! ⚡")
	} else {
		fmt.Println("Status: Tugas singkat. Ayo selesaikan dengan cepat! 🚀")
	}
}
