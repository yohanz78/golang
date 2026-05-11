# About Golang

<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="100" alt="Golang logo">

![Skill Badge - Go](https://img.shields.io/badge/Skill-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Golang, atau sering disebut Go, adalah salah satu bahasa pemrograman yang posisinya sedang sangat strategis di industri saat ini, terutama jika kita berbicara tentang arsitektur backend modern dan komputasi awan (cloud computing).

Berikut adalah gambaran komprehensif mengenai apa yang membuat Golang sangat menonjol:

1. Lahir dari Kebutuhan Industri (Google)

    Go dikembangkan pada tahun 2007 (dan dirilis publik pada 2009) oleh para insinyur senior di Google (Robert Griesemer, Rob Pike, dan Ken Thompson). Bahasa ini diciptakan khusus untuk mengatasi masalah skalabilitas dan kompleksitas pada sistem berskala masif. Mereka menginginkan bahasa yang memiliki performa tinggi seperti C atau C++, tetapi dengan sintaks yang lebih bersih dan waktu kompilasi yang jauh lebih cepat.

2. Fitur Utama dan Kekuatan
    - Sintaks yang Minimalis dan Cepat Dikuasai: Go didesain agar sederhana dan mudah dibaca. Tidak banyak "sihir" atau sintaks tersembunyi. Kesederhanaan ini memungkinkan pengembang untuk mempelajari fondasi inti dengan sangat efisien dan cepat beralih fokus ke pembuatan proyek nyata yang berdampak.

    - Konkurensi (Concurrency) Tingkat Dewa: Ini adalah killer feature dari Go. Go menggunakan apa yang disebut Goroutines. Dibandingkan dengan thread pada sistem operasi tradisional yang memakan banyak memori, Goroutines sangat ringan. Kita bisa menjalankan ribuan bahkan jutaan fungsi secara bersamaan tanpa membebani memori secara drastis.

    - Performa Tinggi (Compiled Language): Go dikompilasi langsung menjadi bahasa mesin (machine code), bukan dijalankan di atas virtual machine atau interpreter. Ini membuatnya sangat cepat saat dieksekusi.

    - Standard Library yang Sangat Kuat: Go memiliki pustaka bawaan (standard library) yang sangat kaya. Misalnya, untuk membuat server web atau membangun RESTful API, kita seringkali tidak membutuhkan framework eksternal pihak ketiga. Cukup menggunakan package bawaan seperti net/http.

3. Ekosistem dan Penggunaan Terpopuler

    Go adalah tulang punggung dari banyak infrastruktur teknologi modern. Teknologi raksasa seperti Docker dan Kubernetes ditulis menggunakan Go. Oleh karena itu, Go sangat dominan di area:
    - Pembuatan Microservices

    - Pengembangan RESTful API dan Backend System

    - Cloud-Native Applications

    - Command-Line Tools (CLI)

4. Nilai Strategis di Pasar Kerja

    Bagi pengembang yang ingin memperkuat pijakan di ranah backend, memasukkan karya berbasis Golang ke dalam sebuah portofolio adalah langkah yang sangat cerdas. Banyak perusahaan, mulai dari startup hingga enterprise, bermigrasi ke arsitektur microservices dan secara aktif mencari talenta yang mampu menangani trafik tinggi dan pemrosesan data asinkron menggunakan Go.

## The Roadmap

Pendekatan _Rapid Skill Acquisition_ sangat relevan untuk skenario ini. Dengan target membangun karya yang solid sebagai portofolio dalam waktu singkat, kita bisa langsung membedah 20% sintaks dan fitur inti Golang yang akan menangani 80% kebutuhan pengembangan _backend_ modern (seperti RESTful API).

Jika sebelumnya sudah terbiasa dengan ekosistem seperti Node.js atau Express, transisi ke Golang akan terasa menyegarkan karena pendekatan _standard library_-nya yang kuat dan penanganan _error_ yang sangat eksplisit.

Berikut adalah _roadmap_ intensif 20 jam yang dibagi menjadi 10 sesi (masing-masing 2 jam).

### Fase 1: Fondasi Sintaks dan Struktur Data (Sesi 1 - 3)

Fokus di sini adalah melatih _muscle memory_ terhadap sintaks Go yang unik dan memahami cara Go menangani data.

**Sesi 1: Setup Lingkungan dan "Hello, World" Steroid (2 Jam)**

- **Materi (20%):** Instalasi Go, inisialisasi Go Modules (`go mod init`), pemahaman fungsi `main()`, deklarasi variabel (`:=` vs `var`), dan tipe data dasar.
- **Praktik (80%):** Membuat program kalkulator sederhana di terminal. Gunakan teknik Feynman: jelaskan setiap baris kode dengan suara lantang untuk memastikan pemahaman logika dasarnya.
- **Sumber Daya:** _A Tour of Go_ (Bagian Basics).

**Sesi 2: Control Flow & Data Collections (2 Jam)**

- **Materi (20%):** `for` loop (Go hanya punya `for`, tidak ada `while`), `if/else`, `switch`, _Arrays_, _Slices_ (sangat penting), dan _Maps_.
- **Praktik (80%):** Membuat program manajemen daftar tugas (_To-Do List_) sederhana berbasis CLI menggunakan _Slices_ dan _Maps_.
- **Sumber Daya:** _Go by Example_ (Slices, Maps, For, If/Else).

**Sesi 3: Structs, Methods, dan Interfaces (2 Jam)**

- **Materi (20%):** Go bukan bahasa PBO (_OOP_) tradisional. Pahami cara membuat tipe data kustom dengan `struct`, menempelkan fungsi pada _struct_ (_methods_), dan menggunakan _interface_ untuk abstraksi.
- **Praktik (80%):** Membuat _struct_ "User" dan "Product", lalu mengimplementasikan _interface_ untuk menghitung diskon atau pajak.
- **Sumber Daya:** Video YouTube "Golang Structs and Interfaces" (channel: _Tech With Tim_ atau _FreeCodeCamp_).

---

### Fase 2: Fitur Unik Go & Performa (Sesi 4 - 6)

Fase ini membahas kekuatan utama Golang: Keamanan memori dan Konkurensi tingkat dewa.

**Sesi 4: Pointers & Error Handling (2 Jam)**

- **Materi (20%):** Konsep memori dasar (Pointers `*` dan `&`) agar hemat RAM. Memahami bahwa _error_ di Go diperlakukan sebagai nilai (_values_), bukan _exception_ (seperti `try/catch`).
- **Praktik (80%):** Memodifikasi kode dari Sesi 2 & 3 agar menggunakan pointer untuk efisiensi memori, dan menambahkan validasi _error_ yang ketat (misal: mencegah pembagian dengan nol).
- **Sumber Daya:** _A Tour of Go_ (Methods and interfaces - mencakup Pointers dan Errors).

**Sesi 5: Dasar Konkurensi - Goroutines & Channels (2 Jam)**

- **Materi (20%):** Menjalankan fungsi di _background_ dengan menambahkan kata `go` di depan fungsi (_Goroutines_). Mengirim data antar _goroutines_ menggunakan _Channels_.
- **Praktik (80%):** Membuat program _scraper_ fiktif yang mengunduh 5 "halaman web" (menggunakan fungsi `time.Sleep` sebagai simulasi penundaan) secara bersamaan, bukan bergantian.
- **Sumber Daya:** _Go by Example_ (Goroutines, Channels).

**Sesi 6: Kontrol Konkurensi - WaitGroup & Mutex (2 Jam)**

- **Materi (20%):** Cara menunggu sekumpulan _goroutines_ selesai bekerja menggunakan `sync.WaitGroup`, dan mencegah manipulasi data yang bentrok (_Race Condition_) dengan `sync.Mutex`.
- **Praktik (80%):** Membangun sistem simulasi loket tiket, di mana beberapa kasir (_goroutines_) mengurangi stok tiket yang sama secara bersamaan dengan aman.
- **Sumber Daya:** Artikel di _DigitalOcean_ (Tutorial berseri Golang Concurrency).

---

### Fase 3: Web & Backend Development (Sesi 7 - 10)

Ini adalah penerapan nyata untuk persiapan masuk ke industri. Go sangat diminati di ranah _backend_ API.

**Sesi 7: Standard Library - HTTP Server & JSON (2 Jam)**

- **Materi (20%):** Menggunakan pustaka bawaan `net/http` untuk membuat _web server_, serta `encoding/json` untuk menerima dan merespons data dalam format JSON.
- **Praktik (80%):** Membuat _server_ lokal yang merespons data JSON statis (misal: profil _user_) saat diakses melalui _browser_ atau _Postman_.
- **Sumber Daya:** Dokumentasi resmi Golang (`pkg.go.dev/net/http`).

**Sesi 8: Routing Lanjutan & Middleware HTTP (2 Jam)**

- **Materi (20%):** Menangani tipe _request_ HTTP (_GET, POST, PUT, DELETE_), membaca URL _parameter_, dan memahami konsep _Middleware_ (misal: mencatat _log_ setiap ada _request_ masuk).
- **Praktik (80%):** Membangun _endpoints_ REST API untuk skenario CRUD (Create, Read, Update, Delete) yang sementara datanya disimpan di memori (_Slices_). Menggunakan _router_ ringan seperti `chi` atau `gorilla/mux` (opsional, `net/http` Go 1.22+ sudah mendukung _routing_ mumpuni).
- **Sumber Daya:** Tutorial YouTube "Golang REST API with Standard Library".

**Sesi 9: Koneksi Database SQL (2 Jam)**

- **Materi (20%):** Menggunakan pustaka `database/sql` beserta _driver_ database (seperti PostgreSQL atau MySQL) untuk menyimpan data secara permanen.
- **Praktik (80%):** Menyambungkan API dari Sesi 8 ke database nyata. Mengganti sistem memori (_Slices_) dengan operasi _query_ SQL `INSERT`, `SELECT`, `UPDATE`, `DELETE`.
- **Sumber Daya:** Tutorial _Go.dev_ bagian "Tutorial: Accessing a relational database".

**Sesi 10: Integrasi Total - Mini Project Portofolio (2 Jam)**

- **Materi (20%):** _Review_ struktur folder yang standar dalam Go (sering disebut _Standard Go Project Layout_) dan _environment variables_ (untuk mengamankan _password_ database).
- **Praktik (80%):** Menyempurnakan API yang sudah dibuat di Sesi 8 & 9. Memastikan penamaan variabel rapi, _error handling_ aman, dan kode berjalan tanpa masalah. Proyek ini siap didorong ke repositori _GitHub_ sebagai pembuktian kemampuan _backend_ yang konkret.
- **Sumber Daya:** _GitHub repo standard-project-layout_ untuk referensi penataan folder.

**Saran Eksekusi:** Gunakan editor _Visual Studio Code_ dengan ekstensi **Go** (dari _Go Team at Google_). Ekstensi ini otomatis melakukan _formatting_ dan membantu mendeteksi _error_ secara seketika (_real-time_). Fokus pada pemahaman kode, bukan sekadar menyalinnya.
