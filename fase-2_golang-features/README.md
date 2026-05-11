**Fase 2: Fitur Unik Go & Performa**

Table of Contents

- [Sesi 4: Pointers \& Error Handling](#sesi-4-pointers--error-handling)
    - [1. Pointers (Manajemen Memori)](#1-pointers-manajemen-memori)
    - [2. Error Handling (Penanganan Masalah)](#2-error-handling-penanganan-masalah)

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
