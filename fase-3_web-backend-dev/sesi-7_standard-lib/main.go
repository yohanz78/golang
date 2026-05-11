package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Mendefinisikan Struct dengan JSON Tags
// Tag json:"..." memberitahu package encoding/json bagaimana cara format outputnya
type Routine struct {
	ID       string `json:"id"`
	Name     string `json:"routine_name"`
	Duration int    `json:"duration_minutes"`
	Priority string `json:"priority_level"`
}

// type Health struct {
// 	Status  string `json:"status"`
// 	Version string `json:"version"`
// }

// Plain response
func HomePage(res http.ResponseWriter, req *http.Request) {
	greeting := "Hello World!"

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(greeting))
}

// 2. Membuat Handler Function untuk Endpoint Routines
func GetRoutinesHandler(w http.ResponseWriter, r *http.Request) {
	// Memastikan endpoint ini hanya menerima method GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Data simulasi (Di dunia nyata, data ini diambil dari Database)
	routines := []Routine{
		{ID: "R-001", Name: "Learn Golang HTTP", Duration: 120, Priority: "High"},
		{ID: "R-002", Name: "Review API Design", Duration: 60, Priority: "Medium"},
		{ID: "R-003", Name: "Read Clean Code", Duration: 30, Priority: "Low"},
	}

	// Memberitahu klien (browser/frontend) bahwa kita mengirim data JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengubah Slice of Structs menjadi JSON dan mengirimkannya ke ResponseWriter (w)
	err := json.NewEncoder(w).Encode(routines)
	if err != nil {
		// Jika terjadi error saat konversi ke JSON, kirim status 500 Internal Server Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ----- Tugas Praktik -----
// Buatlah satu handler function baru bernama HealthCheckHandler.
// 1. Endpoint URL-nya adalah /api/health.
// 2. Jika diakses, endpoint ini akan mengembalikan JSON sederhana: {"status": "API is healthy and running", "version": "1.0.0"}. (Anda bisa membuat Struct baru untuk ini, atau menggunakan map[string]string).
// 3. Daftarkan route tersebut di dalam fungsi main.

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// healthRes := []Health{
	// 	{Status: "API is healthy and running", Version: "1.0.0"},
	// }

	w.Header().Set("Content-Type", "application/json")

	healthRes := map[string]string{
		"status":  "API is healthy and running",
		"version": "1.0.0",
	}

	err := json.NewEncoder(w).Encode(healthRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// 3. Mendaftarkan Route dan Handler-nya
	http.HandleFunc("/", HomePage)

	http.HandleFunc("/api/routines", GetRoutinesHandler)

	http.HandleFunc("/api/health", HealthCheckHandler)

	// Menentukan port server
	port := ":8080"
	fmt.Printf("🚀 Routine Integrator API is running on http://localhost%s\n", port)
	// fmt.Printf("👉 Coba akses: http://localhost%s/api/routines\n", port)

	// 4. Menjalankan Server
	// Nilai 'nil' pada parameter kedua berarti kita menggunakan DefaultServeMux bawaan Golang
	err := http.ListenAndServe(port, nil)

	// Jika server gagal berjalan (misal port 8080 sudah dipakai program lain), program akan berhenti dan mencetak error
	if err != nil {
		fmt.Printf("❌ Server failed to start: %v\n", err)
	}
}
