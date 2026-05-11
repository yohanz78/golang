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
- [Sesi 6: Kontrol Konkurensi - WaitGroup \& Mutex](#sesi-6-kontrol-konkurensi---waitgroup--mutex)
    - [1. WaitGroup (Menunggu Tanpa Menebak)](#1-waitgroup-menunggu-tanpa-menebak)
        - [Contoh Kasus: Absensi Tour Guide di Tempat Wisata](#contoh-kasus-absensi-tour-guide-di-tempat-wisata)
    - [2. Mutex (Mutual Exclusion)](#2-mutex-mutual-exclusion)
        - [Contoh Kasus: Tragedi Dua Orang dan Satu Papan Tulis](#contoh-kasus-tragedi-dua-orang-dan-satu-papan-tulis)
    - [Kesimpulan](#kesimpulan)

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

# Sesi 6: Kontrol Konkurensi - WaitGroup & Mutex

## 1. WaitGroup (Menunggu Tanpa Menebak)

Dari package `sync`. Digunakan untuk menunggu sekumpulan Goroutine selesai mengeksekusi tugasnya. Kita tidak perlu lagi menghitung manual berapa data yang akan dikembalikan melalui _Channel_.

- `wg.Add(1)`: Menambah penghitung tugas sebelum Goroutine dijalankan.
- `wg.Done()`: Mengurangi penghitung (biasanya dipanggil menggunakan `defer` di dalam fungsi Goroutine).
- `wg.Wait()`: Memblokir program utama agar tidak selesai sebelum penghitung kembali menjadi nol.

### Contoh Kasus: Absensi Tour Guide di Tempat Wisata

Sekarang tentang `WaitGroup`. Konsep `WaitGroup` inilah yang paling mirip dengan cara kerja asinkron di JavaScript, lebih tepatnya mirip dengan `Promise.all()`.

Bayangkan `WaitGroup` adalah seorang Tour Guide yang membawa 100 turis (Goroutines) ke sebuah museum.

1. `wg.Add(1)`: Sebelum turis dibiarkan berpencar, Tour Guide mencatat di buku absennya: "Oke, 1 turis pergi. Total yang harus ditunggu: 1." Ini diulang sampai 100.

2. `go ProsesTugas()`: Turis berpencar, mengeksplorasi museum dengan kecepatan masing-masing.

3. `wg.Wait()`: Tour Guide berdiri di depan pintu bus wisata, memblokir bus agar tidak berangkat sebelum buku absennya kembali ke angka 0.

4. `wg.Done()`: Setiap kali seorang turis selesai dan kembali ke bus, ia melapor, "Saya selesai!" dan Tour Guide mengurangi catatannya (-1).

5. Ketika semua 100 turis sudah lapor (`wg.Done()` dipanggil 100 kali), absen kembali ke 0. Barulah `wg.Wait()` selesai memblokir jalan, dan bus (program `main`) boleh berangkat/selesai.

Jika Anda tidak memakai `WaitGroup` di Go, program `main` akan langsung selesai dalam sekejap mata sebelum ke-100 Goroutine sempat menyelesaikan pekerjaannya.

## 2. Mutex (Mutual Exclusion)

Digunakan untuk mencegah **Race Condition**. Saat beberapa Goroutine mencoba membaca dan menulis ke lokasi memori yang sama secara bersamaan, data bisa menjadi korup atau perhitungannya meleset. Mutex "mengunci" variabel tersebut sehingga hanya satu Goroutine yang bisa mengubahnya pada satu waktu.

- `mu.Lock()`: Mengunci akses. Goroutine lain yang mau mengakses harus mengantre.
- `mu.Unlock()`: Membuka kunci agar Goroutine selanjutnya di antrean bisa masuk.

### Contoh Kasus: Tragedi Dua Orang dan Satu Papan Tulis

Bayangkan variabel TotalCompleted adalah sebuah papan tulis di dalam sebuah ruangan.
Goroutines adalah 100 pekerja yang berlarian secara bersamaan (paralel) masuk ke ruangan itu untuk menambahkan angka di papan tulis sebesar +1.

**Tanpa Mutex (Skenario Race Condition):**

1. Pekerja A dan Pekerja B masuk ruangan di milidetik yang sama persis.
2. Keduanya melihat ke papan tulis. Angkanya saat ini: 50.
3. Pekerja A menghitung di kepalanya: 50 + 1 = 51.
4. Pekerja B juga menghitung di kepalanya: 50 + 1 = 51.
5. Pekerja A menghapus angka 50, dan menulis 51.
6. Sepersekian milidetik kemudian, Pekerja B menghapus angka 51 milik A, dan menulis 51 (karena di kepalanya hasil hitungannya tadi adalah 51).

Kedua pekerja sudah menyelesaikan tugasnya, tapi karena mereka membaca dan menulis di waktu yang bertabrakan, hasil akhirnya 51, padahal seharusnya 52! Inilah mengapa hasil Anda 99. Ada dua Goroutine yang saling menimpa (overwrite) data yang sama.

**Dengan Mutex (Mutual Exclusion):**

Mutex adalah sebuah gembok dan kunci di pintu ruangan papan tulis tersebut.

1. Pekerja A sampai duluan. Dia mengunci pintu dari dalam (mu.Lock()).
2. Pekerja B sampai sedetik kemudian. Karena pintu dikunci, dia harus mengantre di luar.
3. Pekerja A membaca 50, mengubahnya jadi 51.
4. Pekerja A keluar dan membuka kunci pintu (mu.Unlock()).
5. Barulah Pekerja B masuk, mengunci pintu (mu.Lock()), membaca angka 51 yang baru, dan mengubahnya menjadi 52.

> **Kenapa di JavaScript tidak ada Mutex?**
>
> Di ekosistem JavaScript (seperti Node.js), arsitekturnya adalah Single-Threaded (Event Loop). Artinya, JS hanya punya satu pekerja yang mengurus papan tulis. Karena pekerjanya cuma satu, tidak akan pernah ada tabrakan! Tapi di Golang, karena kita menggunakan mesin Multi-Threaded secara maksimal, kita punya banyak pekerja yang berjalan bersamaan, sehingga Mutex wajib digunakan untuk melindungi data yang dibagikan (shared data).

## Kesimpulan

- **WaitGroup** memastikan program tidak ditutup sebelum semua tugas latar belakang selesai (Fungsi ekuivalen: `Promise.all()` di JS).

- **Mutex** memastikan data tidak rusak/bertabrakan saat banyak tugas latar belakang mengubah variabel yang sama secara bersamaan (Tidak ada ekuivalen langsung di JS karena JS Single-Threaded).
