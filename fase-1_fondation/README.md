**Fase 1: Fondasi Sintaks dan Struktur Data**

Table of Contents

- [Sesi 1: Setup Lingkungan dan "Hello, World" Steroid](#sesi-1-setup-lingkungan-dan-hello-world-steroid)
  - [Setup Lingkungan](#setup-lingkungan)
  - [Konsep Inti](#konsep-inti)
  - [Deklarasi Variabel (2 Cara Utama)](#deklarasi-variabel-2-cara-utama)
- [Sesi 2: Control Flow \& Data Collections](#sesi-2-control-flow--data-collections)
  - [1. Data Collections (Menyimpan Banyak Data)](#1-data-collections-menyimpan-banyak-data)
  - [2. Control Flow (Alur Logika)](#2-control-flow-alur-logika)
- [Sesi 3: Structs, Methods, dan Interfaces](#sesi-3-structs-methods-dan-interfaces)
  - [1. Structs (Pengganti Class)](#1-structs-pengganti-class)
  - [2. Methods (Fungsi yang Menempel pada Struct)](#2-methods-fungsi-yang-menempel-pada-struct)
  - [3. Interfaces (Kontrak Perilaku)](#3-interfaces-kontrak-perilaku)

# Sesi 1: Setup Lingkungan dan "Hello, World" Steroid

## Setup Lingkungan

Setiap proyek Go modern membutuhkan modul untuk melacak dependensi (mirip `package.json` di Node.js). Jalankan perintah ini di dalam folder proyek:

```bash
go mod init project-name
```

Ini akan menghasilkan file `go.mod`.

## Konsep Inti

1. **Package Main:** Setiap program Go yang bisa dieksekusi harus berada di dalam `package main` dan memiliki fungsi `func main()`. Ini adalah titik awal (_entry point_) program berjalan.
2. **Import:** Keyword `import` digunakan untuk memanggil _library_, seperti `"fmt"` (Format) untuk mencetak teks ke terminal.
3. **Go Mod Init:** Perintah `go mod init <nama-modul>` wajib dijalankan saat memulai proyek baru untuk manajemen dependensi.
4. **`go run main.go`:** Perintah ini akan mengkompilasi dan menjalankan kode untuk folder `main.go`.

## Deklarasi Variabel (2 Cara Utama)

Di Golang, ada dua cara mendeklarasikan variabel:

1. **Eksplisit (menggunakan `var`):** Digunakan jika kita ingin mendeklarasikan variabel tanpa langsung memberinya nilai, atau jika tipe datanya butuh ketegasan.

    `var nama string = "Golang"`

2. **Shorthand / Tipe Inferensi (`:=`):** Ini yang paling sering digunakan (80% kasus). Go akan otomatis menebak tipe datanya berdasarkan nilainya. _Shorthand_ hanya bisa digunakan di dalam fungsi.

    `umur := 22`

# Sesi 2: Control Flow & Data Collections

## 1. Data Collections (Menyimpan Banyak Data)

Di Golang modern, kita jarang menggunakan _Array_ statis. 80% kasus nyata menggunakan **Slices** dan **Maps**.

- **Slices:** Kumpulan data berurutan yang ukurannya bisa bertambah secara dinamis (mirip _List_ di Python atau _Array_ di JavaScript).
  `rutinitas := []string{"Fokus Kerja", "Olahraga"}`
- **Maps:** Kumpulan data berbasis _Key-Value_ (mirip _Object_ di JS atau _Dictionary_ di Python). Sangat cepat untuk mencari data spesifik.
  `durasi := map[string]int{"Fokus Kerja": 120, "Olahraga": 45}`

## 2. Control Flow (Alur Logika)

- **For Loop:** Golang **hanya** memiliki `for` untuk perulangan. Tidak ada `while` atau `do-while`. `for` bisa digunakan untuk mengulang angka, atau mengurai isi _Slice/Map_ menggunakan keyword `range`.
- **If / Else:** Standar pengkondisian logika. Bedanya di Go, kita tidak perlu membungkus kondisinya dengan tanda kurung `()`.
- **Switch:** Sangat efisien di Go. Secara _default_, Go akan langsung berhenti di _case_ yang cocok tanpa perlu menulis `break` di setiap akhir _case_.

# Sesi 3: Structs, Methods, dan Interfaces

## 1. Structs (Pengganti Class)

Struct adalah kumpulan field (variabel) yang dikelompokkan menjadi satu tipe data baru. Ini cara Go merepresentasikan objek.

- **Analogi:** Jika string adalah satu kata, struct adalah sebuah formulir data lengkap.

## 2. Methods (Fungsi yang Menempel pada Struct)

Method adalah fungsi biasa, tetapi ia "menempel" pada sebuah tipe data (biasanya Struct).

- **Ciri Khas:** Memiliki _receiver_ di antara kata `func` dan nama fungsinya.
  `func (t Task) GetDetails() {}` -> `t Task` adalah receiver.

## 3. Interfaces (Kontrak Perilaku)

Interface hanya berisi daftar nama _method_ (tanpa isi/logikanya). Jika sebuah Struct memiliki semua _method_ yang ada di dalam Interface, maka Struct tersebut secara otomatis dianggap mengimplementasikan Interface tersebut (tanpa perlu keyword `implements`).

- **Fungsi Utama:** Membuat kode lebih fleksibel dan modular.
