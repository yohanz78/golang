**Fase 3: Web & Backend Development**

Table of Contents

- [Sesi 7: Standard Library - HTTP Server \& JSON](#sesi-7-standard-library---http-server--json)
    - [1. net/http (Mesin Web Server)](#1-nethttp-mesin-web-server)
    - [2. Handler Function (Penerima Tamu)](#2-handler-function-penerima-tamu)
    - [3. encoding/json (Konversi Data)](#3-encodingjson-konversi-data)
- [Sesi 8: Routing Lanjutan \& Middleware HTTP](#sesi-8-routing-lanjutan--middleware-http)
    - [1. Menangani Method HTTP (GET vs POST)](#1-menangani-method-http-get-vs-post)
    - [2. Decode JSON (Membaca Body Request)](#2-decode-json-membaca-body-request)
    - [3. Middleware HTTP (Penjaga Pintu)](#3-middleware-http-penjaga-pintu)
    - [4. Add-on Explainations: Functions \& Methods on Golang](#4-add-on-explainations-functions--methods-on-golang)
        - [1. Klarifikasi Istilah: "HTTP Method" vs "Golang Method"](#1-klarifikasi-istilah-http-method-vs-golang-method)
        - [2. Apakah aman untuk aplikasi skala besar?](#2-apakah-aman-untuk-aplikasi-skala-besar)
        - [3. Bagaimana Standar Industri Golang Menyelesaikannya?](#3-bagaimana-standar-industri-golang-menyelesaikannya)
- [Sesi 9: Koneksi Database SQL](#sesi-9-koneksi-database-sql)
    - [1. Package `database/sql` \& Driver](#1-package-databasesql--driver)
    - [2. Koneksi: `Open` vs `Ping`](#2-koneksi-open-vs-ping)
    - [3. Eksekusi Query](#3-eksekusi-query)
    - [4. On Production: Environment Variables Setup](#4-on-production-environment-variables-setup)
        - [a. Menginstal Library Pendukung](#a-menginstal-library-pendukung)
        - [b. Membuat File `.env`](#b-membuat-file-env)
        - [c. Membuat File `.gitignore` (Sangat Krusial)](#c-membuat-file-gitignore-sangat-krusial)
        - [d. Modifikasi Kode `main.go`](#d-modifikasi-kode-maingo)
        - [Mengapa Pendekatan Ini Sangat Profesional?](#mengapa-pendekatan-ini-sangat-profesional)

# Sesi 7: Standard Library - HTTP Server & JSON

## 1. net/http (Mesin Web Server)

Golang memiliki package `net/http` yang menyediakan fungsionalitas HTTP client dan server.

- `http.HandleFunc(path, handler)`: Menentukan rute URL (endpoint) dan fungsi apa yang akan dijalankan saat rute tersebut diakses.
- `http.ListenAndServe(port, handler)`: Menjalankan server pada port tertentu (misal: `:8080`). Sifatnya _blocking_, artinya program akan terus berjalan menunggu _request_ (permintaan) masuk.

## 2. Handler Function (Penerima Tamu)

Setiap fungsi _handler_ di Go HTTP wajib memiliki 2 parameter utama:

- `w http.ResponseWriter`: Alat untuk membalas (_Response_) ke klien/browser. (Contoh: mengirim JSON atau status kode 200 OK).
- `r *http.Request`: Pointer yang berisi data permintaan (_Request_) dari klien (Contoh: URL yang diakses, method GET/POST, atau data yang dikirim user).

## 3. encoding/json (Konversi Data)

Untuk mengubah Struct di Go menjadi JSON, kita menggunakan _Struct Tags_.

- Contoh Struct Tag: `` `json:"nama_field"` ``. Ini memberi tahu Go untuk mengubah nama variabel (yang awalnya huruf besar) menjadi format JSON yang kita inginkan (biasanya _camelCase_ atau _snake_case_).
- `json.NewEncoder(w).Encode(data)`: Mengubah data Struct menjadi JSON dan langsung mengirimkannya melalui `http.ResponseWriter`.

# Sesi 8: Routing Lanjutan & Middleware HTTP

## 1. Menangani Method HTTP (GET vs POST)

Satu URL endpoint seringkali memiliki tugas yang berbeda tergantung method-nya.

- `GET`: Meminta/mengambil data.
- `POST`: Mengirim/membuat data baru (biasanya mengirim JSON di dalam Body request).
- Kita bisa menggunakan `if r.Method == http.MethodPost` di dalam satu handler untuk memisahkan logikanya.

## 2. Decode JSON (Membaca Body Request)

Jika `json.NewEncoder(w).Encode(data)` digunakan untuk MENGIRIM JSON ke klien, maka untuk MEMBACA JSON yang dikirim klien ke server kita menggunakan `json.NewDecoder(r.Body).Decode(&data)`.

## 3. Middleware HTTP (Penjaga Pintu)

Middleware adalah fungsi yang berjalan **sebelum** fungsi Handler utama dieksekusi.

- **Analogi:** Middleware adalah satpam di depan gedung. Sebelum tamu (Request) bertemu dengan orang di dalam gedung (Handler), satpam akan mencatat jam kedatangan (Logging), atau mengecek ID tamu (Authentication).
- Di Go, Middleware adalah fungsi yang menerima `http.HandlerFunc` dan mengembalikan `http.HandlerFunc`.

## 4. Add-on Explainations: Functions & Methods on Golang

### 1. Klarifikasi Istilah: "HTTP Method" vs "Golang Method"

Pertama, kita harus membedakan kata "method" dalam konteks ini:

- **HTTP Method:** Ini adalah jenis permintaan dari klien (seperti `GET` untuk mengambil data, `POST` untuk menambah data, `PUT`, `DELETE`).
- **Golang Method:** Ini adalah fungsi yang "menempel" pada sebuah _Struct_ (seperti yang kita pelajari di Sesi 3 dengan _receiver_ `func (t *Task) CompleteTask()`).

Dalam kode Sesi 8 yang baru saja kita buat, kita memasukkan logika untuk mengecek **HTTP Method** (`if r.Method == "GET"` dan `if r.Method == "POST"`) ke dalam satu fungsi utama bernama `RoutinesHandler`.

### 2. Apakah aman untuk aplikasi skala besar?

Jawabannya: **Tidak ideal untuk skala besar.**

Meskipun cara tersebut berfungsi dengan sempurna untuk API skala kecil atau pembelajaran, pendekatan menumpuk banyak `if-else` untuk setiap HTTP Method di dalam satu fungsi akan menjadi mimpi buruk saat aplikasi membesar.

**Analogi Teknik Feynman (Resepsionis Gedung):**
Bayangkan fungsi `RoutinesHandler` adalah seorang Resepsionis di meja lobi.

- Jika perusahaannya kecil (skala kecil), satu resepsionis bisa mengurus tamu yang mau _check-in_ (POST), tamu yang mau bertanya informasi (GET), dan tamu yang mau komplain (DELETE).
- Namun, jika ini adalah perusahaan berskala raksasa dengan ribuan interaksi setiap menit (skala besar), resepsionis tersebut akan kewalahan, antrean menjadi panjang, dan jika dia melakukan satu kesalahan kecil, seluruh operasional lobi berhenti. Kodenya akan memanjang hingga ratusan baris dan sangat sulit di-_maintain_.

### 3. Bagaimana Standar Industri Golang Menyelesaikannya?

Untuk aplikasi berskala besar (terutama saat Anda membangun arsitektur _backend_ untuk dipadukan dengan aplikasi _frontend_ modern berbasis React), standar industrinya adalah **memisahkan satu fungsi untuk satu spesifik HTTP Method**.

Kita biasanya menggunakan _library routing_ pihak ketiga (seperti `go-chi` atau `gorilla/mux`), atau menggunakan pembaruan terbaru di Golang versi 1.22 yang sudah mendukung _routing_ spesifik secara bawaan.

Nantinya, kode Anda tidak akan memakai `if r.Method == ...` lagi, melainkan akan terlihat rapi dan terpisah seperti ini:

```go
// Fungsi khusus menangani GET (Tamu bertanya)
func GetRoutines(w http.ResponseWriter, r *http.Request) {
    // Logika menampilkan data...
}

// Fungsi khusus menangani POST (Tamu check-in)
func CreateRoutine(w http.ResponseWriter, r *http.Request) {
    // Logika menyimpan data...
}

func main() {
    // Standar Golang 1.22 ke atas: Langsung daftarkan HTTP Method-nya di URL
    http.HandleFunc("GET /api/routines", GetRoutines)
    http.HandleFunc("POST /api/routines", CreateRoutine)
}

```

Dengan memecahnya seperti ini, jika terjadi _error_ pada saat menyimpan data baru (POST), fungsi untuk menampilkan data (GET) tidak akan terpengaruh sama sekali. Kode menjadi lebih bersih (_clean code_) dan terisolasi dengan aman.

# Sesi 9: Koneksi Database SQL

## 1. Package `database/sql` & Driver

Golang menyediakan package bawaan `database/sql` yang sangat tangguh. Namun, package ini hanya bertindak sebagai "antarmuka" (_interface_) universal. Untuk benar-benar terhubung ke _database_ spesifik (seperti PostgreSQL, MySQL, atau SQLite), kita wajib mengunduh **Driver** eksternal.

- Contoh driver PostgreSQL: `github.com/lib/pq`

## 2. Koneksi: `Open` vs `Ping`

- `sql.Open()`: Hanya memvalidasi format _connection string_ (URL kredensial database), tetapi **belum** benar-benar membuka koneksi fisik ke jaringan _database_.
- `db.Ping()`: Digunakan segera setelah `Open()` untuk memaksa Go mengetuk pintu _database_ dan memastikan koneksi benar-benar berhasil terjalin.

## 3. Eksekusi Query

- `db.Exec()`: Digunakan untuk operasi yang tidak mengembalikan baris data (seperti `INSERT`, `UPDATE`, `DELETE`, atau `CREATE TABLE`).
- `db.Query()`: Digunakan untuk mengambil banyak baris data (`SELECT`).
- `db.QueryRow()`: Digunakan untuk mengambil tepat satu baris data saja.

**Catatan Penting Penulisan Query SQL:** Saat Anda mulai merancang _query_ lanjutan di dalam Go yang melibatkan penggabungan tabel, ingat kembali aturan logika struktural kita untuk mencegah kesalahan penarikan data: eksekusi `INNER JOIN` akan melibatkan semua data walaupun `departemen_id`-nya masih kosong. Sedangkan `LEFT JOIN`, hanya akan menampilkan `departemen_id` yang tidak kosong.

## 4. On Production: Environment Variables Setup

Di dunia profesional, menyimpan kredensial (seperti _password database_, _API Keys_, atau _Secret Token_) langsung di dalam kode sumber (_hardcode_) adalah sebuah risiko keamanan yang sangat fatal. Jika kode tersebut diunggah ke GitHub, siapa saja bisa melihat _password database_ Anda.

Standar industri mutlak untuk menangani ini adalah menggunakan **Environment Variables** (Variabel Lingkungan), biasanya disimpan dalam sebuah file bernama `.env`.

Berikut adalah cara memodifikasi proyek Anda agar memenuhi standar keamanan profesional.

### a. Menginstal Library Pendukung

Walaupun Golang bisa membaca _Environment Variables_ bawaan sistem operasi, untuk tahap pengembangan lokal (_local development_), kita menggunakan _library_ pihak ketiga yang sangat populer bernama `godotenv` untuk membaca file `.env`.

Jalankan perintah ini di terminal:

```bash
go get github.com/joho/godotenv

```

### b. Membuat File `.env`

Buat sebuah file baru tepat di folder yang sama dengan `main.go`, dan beri nama `.env` (tanpa nama depan, hanya ekstensi). Isi dengan kredensial Anda:

```env
# File: .env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=rahasia
DB_NAME=routine_db
DB_SSLMODE=disable

```

### c. Membuat File `.gitignore` (Sangat Krusial)

Untuk mencegah file `.env` yang berisi _password_ ikut terunggah ke repositori seperti GitHub, buat file bernama `.gitignore` di folder utama proyek, dan isi dengan:

```text
# File: .gitignore
.env

```

### d. Modifikasi Kode `main.go`

Sekarang, kita perbarui kode kita untuk memuat file `.env` menggunakan _library_ `godotenv`, lalu mengambil nilainya menggunakan _package_ bawaan Golang `os`.

```go
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

```

### Mengapa Pendekatan Ini Sangat Profesional?

1. **Keamanan Ekstra:** _Password_ tidak pernah tercatat di dalam riwayat (_history_) kode (Git).
2. **Fleksibilitas (_Portability_):** Jika aplikasi Anda dijalankan di komputer _programmer_ lain atau di _server cloud_ (seperti AWS atau Heroku), kita tidak perlu merombak _file_ `main.go`. Cukup ubah saja nilai yang ada di _file_ `.env` atau _environment settings_ di _server_.
3. **Kode Bersih (_Clean Code_):** Memisahkan fungsi `LoadConfig()` membuat alur utama (`main()`) tetap fokus pada proses bisnis, bukan tenggelam dalam logika pembacaan sistem (_Separation of Concerns_).
