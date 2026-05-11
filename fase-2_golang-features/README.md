**Fase 2: Fitur Unik Go & Performa**

Table of Contents

- [Sesi 4: Pointers \& Error Handling](#sesi-4-pointers--error-handling)
    - [1. Pointers (Manajemen Memori)](#1-pointers-manajemen-memori)
    - [2. Error Handling (Penanganan Masalah)](#2-error-handling-penanganan-masalah)
- [Sesi 5: Dasar Konkurensi - Goroutines \& Channels](#sesi-5-dasar-konkurensi---goroutines--channels)
    - [1. Goroutines (Lightweight Threads)](#1-goroutines-lightweight-threads)
    - [2. Channels (Pipa Komunikasi)](#2-channels-pipa-komunikasi)
- [Sesi 5: Dasar Konkurensi - Goroutines \& Channels](#sesi-5-dasar-konkurensi---goroutines--channels-1)
    - [1. Goroutines (Lightweight Threads)](#1-goroutines-lightweight-threads-1)
    - [2. Channels (Pipa Komunikasi)](#2-channels-pipa-komunikasi-1)

# Sesi 4: Pointers & Error Handling

## 1. Pointers (Manajemen Memori)

Di Go, secara _default_ saat kita mempassing data (seperti Struct) ke dalam sebuah fungsi, Go akan membuat **salinan (copy)** dari data tersebut. Jika datanya besar, ini boros memori.
Pointer adalah variabel yang menyimpan _alamat memori_ dari variabel lain, bukan menyimpan nilainya secara langsung.

- `&` (Ampersand): Digunakan untuk mengambil/melihat alamat memori dari sebuah variabel.
- `*` (Asterisk): Digunakan untuk mendeklarasikan tipe pointer, atau untuk mengakses/mengubah nilai asli dari sebuah alamat memori (_dereferencing_).
- **Aturan Emas:** Gunakan Pointer saat Anda ingin **memodifikasi nilai asli** dari sebuah data di dalam fungsi, atau saat mempassing _Struct_ berukuran sangat besar.

## 2. Error Handling (Penanganan Masalah)

Go tidak memiliki blok `try / catch` / `finally` seperti bahasa lain. Di Go, _error_ diperlakukan sebagai nilai (value) biasa yang direturn oleh fungsi.

- Fungsi di Go biasanya mengembalikan multi-nilai (contoh: hasil data dan error).
- Pola standar yang wajib ada: Selalu periksa apakah `err != nil`. Jika tidak `nil`, berarti fungsi tersebut mengalami kegagalan.

# Sesi 5: Dasar Konkurensi - Goroutines & Channels

## 1. Goroutines (Lightweight Threads)

Goroutine adalah fungsi yang dieksekusi secara independen (konkuren) dan dikelola langsung oleh _Go Runtime_, bukan sistem operasi. Goroutine sangat ringan (hanya memakan memori sekitar 2KB saat awal dibuat).

- **Sintaks:** Cukup tambahkan _keyword_ `go` sebelum pemanggilan fungsi.
  `go FetchData()`

## 2. Channels (Pipa Komunikasi)

Karena Goroutine berjalan secara terpisah di latar belakang, kita butuh cara agar mereka bisa saling bertukar data dengan aman dan memberi tahu fungsi `main` bahwa tugas mereka sudah selesai.

- **Membuat Channel:** `ch := make(chan string)` -> Membuat channel yang hanya bisa mengalirkan data _string_.
- **Mengirim Data ke Channel:** `ch <- "Data Selesai"` -> Menggunakan operator panah kiri.
- **Menerima Data dari Channel:** `hasil := <-ch` -> Proses ini bersifat _blocking_. Artinya, program akan berhenti sementara di baris ini sampai ada data yang masuk ke channel.

# Sesi 5: Dasar Konkurensi - Goroutines & Channels

## 1. Goroutines (Lightweight Threads)

Goroutine adalah fungsi yang dieksekusi secara independen (konkuren) dan dikelola langsung oleh _Go Runtime_, bukan sistem operasi. Goroutine sangat ringan (hanya memakan memori sekitar 2KB saat awal dibuat).

- **Sintaks:** Cukup tambahkan _keyword_ `go` sebelum pemanggilan fungsi.
  `go FetchData()`

## 2. Channels (Pipa Komunikasi)

Karena Goroutine berjalan secara terpisah di latar belakang, kita butuh cara agar mereka bisa saling bertukar data dengan aman dan memberi tahu fungsi `main` bahwa tugas mereka sudah selesai.

- **Membuat Channel:** `ch := make(chan string)` -> Membuat channel yang hanya bisa mengalirkan data _string_.
- **Mengirim Data ke Channel:** `ch <- "Data Selesai"` -> Menggunakan operator panah kiri.
- **Menerima Data dari Channel:** `hasil := <-ch` -> Proses ini bersifat _blocking_. Artinya, program akan berhenti sementara di baris ini sampai ada data yang masuk ke channel.
