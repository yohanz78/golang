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
