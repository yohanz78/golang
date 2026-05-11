package main

import (
	"errors" // Package bawaan untuk membuat error kustom
	"fmt"
)

// 1. Struct Task dengan tambahan status
type Task struct {
	Name        string
	Duration    int
	Priority    string
	IsCompleted bool
}

// 2. Method TANPA Pointer (Pass by Value)
// Method ini HANYA akan memodifikasi data copy-an, TIDAK akan mengubah data asli
func (t Task) RenameTask(newName string) {
	t.Name = newName
}

// 3. Method DENGAN Pointer (Pass by Reference)
// Menggunakan *Task agar kita memodifikasi data langsung di alamat memori aslinya
func (t *Task) ExtendDuration(extraMinutes int) {
	t.Duration += extraMinutes // Nilai asli akan berubah
}

// --- TUGAS PRAKTIK ---
// Method baru dengan Pointer Receiver untuk mengubah status IsCompleted
func (t *Task) CompleteTask() {
	// Karena menggunakan *, nilai di memori asli akan diubah menjadi true
	t.IsCompleted = true
}

// 4. Error Handling pada Fungsi
// Fungsi ini mengembalikan 2 nilai: Pointer dari Task (*Task) dan sebuah error
func CreateTask(name string, duration int, priority string) (*Task, error) {
	// Validasi: Durasi tidak masuk akal
	if duration <= 0 {
		// Return nil untuk data Task, dan return pesan error-nya
		return nil, errors.New("invalid duration: must be greater than 0")
	}

	// Validasi: Nama kosong
	if name == "" {
		return nil, errors.New("invalid name: task name cannot be empty")
	}

	// Jika lolos validasi, buat task baru
	newTask := Task{
		Name:        name,
		Duration:    duration,
		Priority:    priority,
		IsCompleted: false,
	}

	// Mengembalikan alamat memori dari newTask (&) dan error nil (berarti sukses/tidak ada error)
	return &newTask, nil
}

func main() {
	fmt.Println("=== Routine Integrator: Pointers & Errors ===")

	// --- DEMO POINTERS ---
	fmt.Println("--- 1. Pointers Demo ---")
	myTask := Task{Name: "Learn Pointers", Duration: 60, Priority: "High", IsCompleted: false}
	fmt.Printf("Data Awal: %+v\n", myTask)

	// Mencoba mengubah nama tanpa pointer (Gagal mengubah nilai asli)
	myTask.RenameTask("Mastering Pointers")
	fmt.Printf("Setelah RenameTask (Tanpa Pointer): %+v\n", myTask)

	// Mengubah durasi dengan pointer (Berhasil mengubah nilai asli)
	myTask.ExtendDuration(30)
	fmt.Printf("Setelah ExtendDuration (Dengan Pointer): %+v\n\n", myTask)

	// Mengubah IsCompleted dengan pointer
	myTask.CompleteTask()
	fmt.Printf("Setelah CompleteTask (Dengan Pointer): %+v\n\n", myTask)

	// --- DEMO ERROR HANDLING ---
	fmt.Println("--- 2. Error Handling Demo ---")

	// Skenario A: Input Valid
	validTask, err := CreateTask("Setup Database", 120, "High")
	if err != nil { // Menangkap dan mengecek error
		fmt.Printf("❌ Gagal membuat tugas: %v\n", err)
	} else {
		fmt.Printf("✅ Sukses membuat tugas: %s (%d menit)\n", validTask.Name, validTask.Duration)
	}

	// Skenario B: Input Tidak Valid (Durasi negatif)
	invalidTask, err := CreateTask("Tidur Siang", -10, "Low")
	if err != nil { // Menangkap error
		fmt.Printf("❌ Gagal membuat tugas: %v\n", err)
	} else {
		fmt.Printf("✅ Sukses membuat tugas: %s\n", invalidTask.Name)
	}
}
