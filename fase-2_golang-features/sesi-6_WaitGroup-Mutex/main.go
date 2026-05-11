package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct untuk melacak total tugas yang diselesaikan
type TaskTracker struct {
	TotalCompleted int
	// Menambahkan Mutex ke dalam struct untuk melindungi properti TotalCompleted
	Mu sync.Mutex
}

// Method untuk menyelesaikan tugas (dijalankan sebagai Goroutine)
// Kita WAJIB mempassing WaitGroup menggunakan Pointer (*sync.WaitGroup)
func (t *TaskTracker) ProcessTask(taskID int, wg *sync.WaitGroup) {
	// Defer memastikan wg.Done() pasti dipanggil di akhir fungsi,
	// meskipun nanti ada error di tengah-tengah fungsi.
	defer wg.Done()

	// Mensimulasikan waktu pemrosesan tugas
	time.Sleep(10 * time.Millisecond)

	// --- CRITICAL SECTION MULAI ---
	// Mengunci akses memori. Goroutine lain harus menunggu.
	t.Mu.Lock()

	t.TotalCompleted++ // Mengubah shared data

	// Membuka kunci memori agar Goroutine lain bisa mengubah data
	t.Mu.Unlock()
	// --- CRITICAL SECTION SELESAI ---

	fmt.Printf("Task %d processed.\n", taskID)
}

func main() {
	fmt.Println("=== Routine Integrator: WaitGroup & Mutex ===\n")

	// 1. Inisialisasi struct TaskTracker
	tracker := TaskTracker{TotalCompleted: 0}

	// 2. Inisialisasi WaitGroup
	var wg sync.WaitGroup

	totalTasks := 100
	fmt.Printf("Memulai pemrosesan %d tugas secara konkuren...\n", totalTasks)

	// Menjalankan 100 Goroutines
	for i := 1; i <= totalTasks; i++ {
		// Menambah antrean di WaitGroup SETIAP KALI sebelum Goroutine berjalan
		wg.Add(1)

		// Menjalankan method sebagai Goroutine
		go tracker.ProcessTask(i, &wg)
	}

	// 3. Program utama berhenti sejenak di sini, menunggu antrean wg menjadi 0
	wg.Wait()

	// Mencetak hasil akhir setelah semua Goroutine selesai
	fmt.Printf("\nSemua tugas selesai diproses!\n")
	fmt.Printf("Total Tugas Selesai (Target 100): %d\n", tracker.TotalCompleted)
}
