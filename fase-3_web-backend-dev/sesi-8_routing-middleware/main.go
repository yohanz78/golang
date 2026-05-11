package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Struct untuk data tugas
type Routine struct {
	ID       string `json:"id"`
	Name     string `json:"routine_name"`
	Duration int    `json:"duration_minutes"`
}

// Simulasi Database di memori (menggunakan Slice)
var dbRoutines = []Routine{
	{ID: "R-001", Name: "Learn Golang HTTP", Duration: 120},
}

// 1. Handler Utama (Menangani GET dan POST)
func RoutinesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Jika Method GET: Tampilkan semua data
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(dbRoutines)
		return
	}

	// Jika Method POST: Tambahkan data baru
	if r.Method == http.MethodPost {
		var newRoutine Routine

		// Membaca (Decode) JSON dari request body yang dikirim oleh user
		// Perhatikan penggunaan pointer (&newRoutine) agar nilainya langsung diubah
		err := json.NewDecoder(r.Body).Decode(&newRoutine)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// Menambahkan data baru ke dalam "database" slice kita
		dbRoutines = append(dbRoutines, newRoutine)

		// Mengirimkan status 201 Created dan mengembalikan data yang baru dibuat
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newRoutine)
		return
	}

	// Jika method bukan GET atau POST
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// 2. Middleware Function (Satpam Penjaga Pintu)
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Kode di sini dieksekusi SEBELUM masuk ke handler utama
		fmt.Printf("Menerima request: [%s] %s\n", r.Method, r.URL.Path)

		// Mempersilakan request masuk ke handler utama (RoutinesHandler)
		next(w, r)

		// Kode di sini dieksekusi SETELAH handler utama selesai
		duration := time.Since(start)
		fmt.Printf("Selesai diproses dalam: %v\n", duration)
	}
}

func main() {
	fmt.Println("=== Routine Integrator API dengan Middleware ===")

	// 3. Mendaftarkan Route dengan Middleware
	// RoutinesHandler dibungkus (wrapped) oleh LoggerMiddleware
	http.HandleFunc("/api/routines", LoggerMiddleware(RoutinesHandler))

	port := ":8080"
	fmt.Printf("🚀 Server berjalan di port http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
