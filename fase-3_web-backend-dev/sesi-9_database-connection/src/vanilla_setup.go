package vanilla_setup

import (
	"database/sql"
	"fmt"
	"log"

	// Underscore (_) di depan berarti kita melakukan "Blank Identifier"
	// Kita hanya memicu fungsi init() dari driver ini agar mendaftarkan dirinya
	// ke package database/sql, tanpa memanggil fungsinya secara langsung.
	_ "github.com/lib/pq"
)

// Struct untuk menampung hasil pemetaan data dari database
type Routine struct {
	ID       int
	Name     string
	Duration int
}

func vanilla_setup() {
	fmt.Println("=== Routine Integrator: Database SQL ===\n")

	// 1. Konfigurasi Koneksi (Ubah sesuai dengan kredensial PostgreSQL lokal Anda)
	// Format: "host=... port=... user=... password=... dbname=... sslmode=disable"
	connStr := "user=postgres password=DBadmin789 dbname=golang_test sslmode=disable"

	// 2. Membuka Koneksi
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Gagal inisialisasi koneksi: %v", err)
	}
	// Memastikan koneksi ditutup secara otomatis tepat sebelum fungsi main berakhir
	defer db.Close()

	// 3. Memastikan Koneksi Fisik Berhasil
	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Gagal terhubung. Pastikan PostgreSQL menyala dan kredensial benar: %v", err)
	}
	fmt.Println("✅ Berhasil terhubung ke database PostgreSQL!")

	// 4. Eksekusi DDL: Membuat Tabel jika belum ada
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS routines (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		duration INT NOT NULL
	);`

	// Menggunakan Exec karena tidak mengharapkan kembalian baris data
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("❌ Gagal membuat tabel: %v", err)
	}
	fmt.Println("✅ Tabel 'routines' siap digunakan.")

	// 5. Menyisipkan Data Baru (INSERT)
	// $1 dan $2 adalah placeholder untuk mencegah SQL Injection
	insertQuery := `INSERT INTO routines (name, duration) VALUES ($1, $2) RETURNING id`
	var newID int

	// Menggunakan QueryRow karena ada perintah 'RETURNING id' (mengharapkan 1 data kembali)
	err = db.QueryRow(insertQuery, "Belajar Golang SQL", 90).Scan(&newID)
	if err != nil {
		log.Fatalf("❌ Gagal menyisipkan data: %v", err)
	}
	fmt.Printf("✅ Berhasil menambahkan tugas baru dengan ID Database: %d\n", newID)

	// 6. Mengambil Data (SELECT)
	fmt.Println("\n--- Daftar Tugas dari Database ---")

	// Menggunakan Query karena mengharapkan banyak baris data
	rows, err := db.Query("SELECT id, name, duration FROM routines")
	if err != nil {
		log.Fatalf("❌ Gagal mengambil data: %v", err)
	}
	// Jangan lupa menutup objek rows untuk mencegah memory leak
	defer rows.Close()

	// Looping untuk membaca setiap baris yang dikembalikan database
	for rows.Next() {
		var r Routine

		// Scan berfungsi memindahkan data dari kolom database ke dalam variabel Struct
		// Urutan Scan harus sama persis dengan urutan SELECT di atas
		err := rows.Scan(&r.ID, &r.Name, &r.Duration)
		if err != nil {
			log.Println("Gagal membaca baris data:", err)
			continue
		}

		fmt.Printf("ID: %d | Tugas: %s | Durasi: %d menit\n", r.ID, r.Name, r.Duration)
	}
}
