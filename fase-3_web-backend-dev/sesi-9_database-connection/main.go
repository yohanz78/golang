package main

import (
	"database/sql"
	"fmt"
	"log"
	"os" // Package bawaan untuk mengakses sistem operasi / environment variables

	"github.com/joho/godotenv" // Package untuk memuat file .env
	_ "github.com/lib/pq"
)

// Struct untuk menampung hasil pemetaan data
type Routine struct {
	ID       int
	Name     string
	Duration int
}

// Fungsi khusus untuk memuat konfigurasi
// Ini memisahkan logika konfigurasi dari logika utama (Separation of Concerns)
func LoadConfig() {
	// Memuat file .env yang ada di direktori yang sama
	err := godotenv.Load()
	if err != nil {
		// Kita menggunakan log.Println, bukan Fatal, karena di server Production (seperti AWS/GCP),
		// file .env biasanya tidak ada. Variabel disuntikkan langsung ke server.
		log.Println("⚠️ File .env tidak ditemukan, menggunakan variabel environment bawaan sistem.")
	}
}

func main() {
	fmt.Println("=== Routine Integrator: Secure Database Connection ===\n")

	// 1. Memuat konfigurasi dari .env
	LoadConfig()

	// 2. Mengambil nilai dari Environment Variables
	// os.Getenv("NAMA_KEY") akan membaca nilai dari file .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// 3. Merakit Connection String (connStr) secara dinamis
	// Sprintf digunakan untuk memformat string tanpa langsung mencetaknya ke terminal
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// 4. Membuka Koneksi (Proses selanjutnya sama persis seperti sebelumnya)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Gagal inisialisasi koneksi: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Gagal terhubung ke database: %v", err)
	}
	fmt.Println("✅ Berhasil terhubung ke database PostgreSQL secara aman!")

	// ... (Sisa kode untuk CREATE TABLE, INSERT, dan SELECT diletakkan di sini) ...
}
