package main

import (
	"fmt"
	"time"
)

// Fungsi ini mensimulasikan proses penarikan data dari API yang memakan waktu lambat
func FetchRoutineInspiration(source string, resultChan chan string) {
	// Mensimulasikan delay jaringan selama 2 detik
	time.Sleep(2 * time.Second)

	// Membuat format pesan hasil data yang didapat
	result := fmt.Sprintf("Inspiration from %s: Try Pomodoro Technique", source)

	// Mengirim hasil string tersebut ke dalam channel
	resultChan <- result
}

func main() {
	fmt.Println("=== Routine Integrator: Concurrency (Goroutines) ===\n")

	// Mencatat waktu mulai untuk melihat kecepatan eksekusi
	startTime := time.Now()

	// 1. Membuat channel yang akan mengalirkan data bertipe string
	inspirationChan := make(chan string)

	// Daftar sumber data tiruan
	sources := []string{"Google Calendar", "Notion API", "Medium Articles", "Twitter Trends", "YouTube API"}

	fmt.Println("Memulai proses penarikan data (Fetching)...")

	// 2. Looping untuk menjalankan Goroutines
	for _, src := range sources {
		// Menjalankan fungsi secara KUREN / PARALEL dengan keyword 'go'
		// Setiap iterasi loop tidak akan menunggu fungsi sebelumnya selesai
		go FetchRoutineInspiration(src, inspirationChan)
	}

	// 3. Menerima data dari channel
	// Karena kita menjalankan 3 Goroutine, kita juga harus menunggu/menerima 3 balasan dari channel
	for i := 0; i < len(sources); i++ {
		// Program akan 'menunggu' (blocking) di baris ini sampai ada data yang masuk ke channel
		data := <-inspirationChan
		fmt.Printf("✅ Berhasil menerima: %s\n", data)
	}

	// Menghitung total waktu eksekusi
	duration := time.Since(startTime)
	fmt.Printf("\nSelesai! Total waktu eksekusi: %v\n", duration)

	// Catatan Penting:
	// Walaupun ada 3 proses yang masing-masing butuh 2 detik,
	// total waktunya akan tetap sekitar 2 detik (bukan 6 detik), karena ketiganya berjalan bersamaan!
}
