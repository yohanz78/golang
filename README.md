# About Golang

<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="100" alt="Golang logo">

![Skill Badge - Go](https://img.shields.io/badge/Skill-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

**Table of Contents**

- [About Golang](#about-golang)
    - [The Roadmap](#the-roadmap)
        - [Fase 1: Fondasi Sintaks dan Struktur Data (Sesi 1 - 3)](#fase-1-fondasi-sintaks-dan-struktur-data-sesi-1---3)
        - [Fase 2: Fitur Unik Go \& Performa (Sesi 4 - 6)](#fase-2-fitur-unik-go--performa-sesi-4---6)
        - [Fase 3: Web \& Backend Development (Sesi 7 - 10)](#fase-3-web--backend-development-sesi-7---10)
- [QNA](#qna)
    - [API Architectures: RESTful API and GraphQL](#api-architectures-restful-api-and-graphql)
        - [1. RESTful API](#1-restful-api)
        - [2. GraphQL](#2-graphql)
    - [Software Architectures: From Small to Large Scale](#software-architectures-from-small-to-large-scale)
        - [1. Skala Kecil: Flat / Layered Structure (MVC-lite)](#1-skala-kecil-flat--layered-structure-mvc-lite)
        - [2. Skala Menengah - Besar: Standard Go Project Layout](#2-skala-menengah---besar-standard-go-project-layout)
        - [3. Skala Enterprise: Clean Architecture / Domain-Driven Design (DDD)](#3-skala-enterprise-clean-architecture--domain-driven-design-ddd)
        - [Strategi Eksekusi Portofolio](#strategi-eksekusi-portofolio)

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

# QNA

## API Architectures: RESTful API and GraphQL

Jawaban singkatnya: **Sangat mendukung keduanya.** Golang adalah salah satu bahasa pilihan utama di industri saat ini, baik untuk membangun RESTful API maupun GraphQL.

Berikut adalah pembedahannya dari kacamata arsitektur _backend_ Go:

### 1. RESTful API

REST adalah "warga negara kelas satu" (_first-class citizen_) di ekosistem Go. Seperti yang sudah kita praktikkan, Anda bahkan tidak memerlukan pustaka eksternal untuk membuat REST API berkinerja tinggi berkat paket bawaan `net/http`.

Namun, untuk pengembangan skala besar, ekosistem Go menawarkan berbagai pilihan yang lebih canggih:

- **Router Ringan:** Banyak perusahaan menggunakan _router_ pihak ketiga seperti `chi` atau `gorilla/mux`. _Library_ ini mempermudah penulisan rute URL (seperti memisahkan GET dan POST) tanpa mengubah gaya struktur asli dari `net/http`.
- **Web Framework (Performa Ekstrem):** Jika Anda sudah memiliki basis pemahaman di ekosistem JavaScript (seperti Node/Express), Go memiliki _framework_ modern bernama **Fiber**. _Framework_ ini sangat populer karena sintaksnya terinspirasi dari Express.js dan dirancang untuk performa ekstrem. Fiber sangat mumpuni untuk menangani arsitektur platform yang kompleks dengan trafik tinggi, seperti sistem manajemen acara (_event_) dan integrasi media.

### 2. GraphQL

Golang juga memiliki dukungan komunitas yang sangat kuat untuk arsitektur GraphQL. Meskipun tidak ada dukungan langsung di _standard library_-nya, alat-alat pihak ketiga yang tersedia sudah memiliki kualitas _enterprise_.

Di industri profesional, ada satu _library_ yang menjadi standar _de facto_ untuk GraphQL di Go, yaitu **`gqlgen`** (GraphQL Go Generate).

- **Pendekatan Schema-First:** Anda mendefinisikan tipe data dan _query_ GraphQL Anda di sebuah file skema khusus (`.graphqls`).
- **Code Generation (Sihirnya Go):** `gqlgen` akan membaca file skema tersebut dan **secara otomatis menuliskan kode Golang** (_structs_ dan _interfaces_) untuk Anda. Tugas Anda hanya tinggal mengisi logika bisnis penarikan datanya (_resolvers_).
- **Type-Safe:** Pendekatan ini menjamin keakuratan tipe data antara apa yang diminta oleh klien (aplikasi _frontend_ React) dan apa yang dikirim oleh server, sehingga meminimalisir _bug_ saat peluncuran.

**Kesimpulan Penggunaan:**
Arsitektur Go sangat fleksibel. Jika aplikasi Anda membutuhkan pertukaran data yang sangat dinamis dari sisi _frontend_ agar tidak terjadi _over-fetching_ (menarik data berlebih), GraphQL menggunakan `gqlgen` adalah solusi yang sangat elegan. Namun, untuk fondasi _microservices_ yang membutuhkan kecepatan mentah, REST API dengan Fiber atau paket bawaan Go tetap menjadi primadona.

## Software Architectures: From Small to Large Scale

Ini adalah topik yang sangat krusial. Memahami arsitektur sejak awal adalah investasi 20% usaha yang akan mencegah 80% rasa frustrasi saat proyek Anda mulai membesar dan kompleks.

Golang memiliki filosofi yang unik: ia sangat ketat soal format penulisan kode, tetapi **sangat bebas** (_unopinionated_) soal struktur folder. Tidak ada aturan baku dari pembuat bahasanya. Namun, seiring berjalannya waktu, komunitas dan industri telah membentuk pola standar.

Berikut adalah evolusi arsitektur dan struktur folder di Golang, dari aplikasi skala kecil hingga _enterprise_.

---

### 1. Skala Kecil: Flat / Layered Structure (MVC-lite)

Untuk aplikasi _microservices_ tunggal atau proyek awal seperti _Routine Integrator_ versi dasar, arsitektur yang terlalu kompleks justru akan memperlambat Anda (_overengineering_). Struktur ini mirip dengan konsep Model-View-Controller (MVC) klasik.

**Konsep:** Memisahkan kode berdasarkan _fungsi teknisnya_ (semua _handler_ kumpul jadi satu, semua koneksi _database_ kumpul jadi satu).

**Struktur Folder:**

```text
my-app/
├── main.go           # Titik awal aplikasi berjalan
├── .env              # Variabel konfigurasi rahasia
├── handlers/         # (Atau controllers) Berisi logika HTTP GET, POST, dst
│   └── routine.go
├── models/           # Berisi definisi Struct (seperti struct Routine)
│   └── routine.go
├── database/         # Konfigurasi koneksi SQL
│   └── postgres.go
└── go.mod

```

---

### 2. Skala Menengah - Besar: Standard Go Project Layout

Ketika aplikasi mulai memiliki banyak fitur (misalnya, Anda sedang merancang sebuah _Event & Media Management Platform_ yang menggunakan _framework_ performa tinggi seperti Fiber), struktur _flat_ akan membuat folder `handlers` Anda membengkak dan sulit dibaca.

Industri biasanya mengadopsi repositori **`golang-standards/project-layout`**. Ini adalah struktur yang paling umum Anda temui di _repository_ perusahaan.

**Konsep:** Memisahkan mana kode yang bersifat publik (_bisa di-import oleh proyek Go lain_) dan mana yang bersifat privat (_hanya untuk internal aplikasi ini_), serta memisahkan titik masuk aplikasi (`cmd`).

**Struktur Folder:**

```text
event-platform/
├── cmd/
│   └── api/
│       └── main.go       # Titik awal. Hanya berisi inisialisasi server, DB, dan router.
├── internal/             # KODE PRIVAT. Fitur utama aplikasi Anda hidup di sini.
│   ├── events/           # Dikelompokkan berdasarkan FITUR (Domain), bukan teknis
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   └── media/
│       ├── handler.go
│       └── service.go
├── pkg/                  # KODE PUBLIK. Utility yang bisa dipakai proyek lain (misal: logger kustom, format helper)
├── docs/                 # Dokumentasi API (Swagger / OpenAPI)
├── .env
└── go.mod

```

---

### 3. Skala Enterprise: Clean Architecture / Domain-Driven Design (DDD)

Untuk sistem berskala raksasa di mana tim _frontend_ dan _backend_ bekerja terpisah, dan bisnis menuntut perubahan _database_ atau _framework_ tanpa merombak ulang sistem, industri menggunakan **Clean Architecture** (dipopulerkan oleh Uncle Bob).

**Konsep Utama (Separation of Concerns):** Aturan mutlaknya adalah _Dependency Rule_. Lapisan luar (Web/Database) bergantung pada lapisan dalam (Logika Bisnis), **bukan** sebaliknya. Logika inti bisnis tidak boleh tahu sama sekali apakah Anda memakai PostgreSQL atau MySQL, atau apakah Anda memakai _framework_ Fiber, standard `net/http`, atau bahkan gRPC.

**Struktur Folder (Pendekatan Modular):**

```text
enterprise-app/
├── cmd/api/main.go
├── internal/
│   └── routine/
│       ├── domain.go       # (Layer 1) Definisi Struct (Entities) dan Interface (Kontrak mutlak)
│       ├── repository/     # (Layer 2) Berkomunikasi dengan Database (implementasi SQL)
│       ├── usecase/        # (Layer 3) Logika Bisnis murni. Mengatur alur data, validasi, perhitungan.
│       └── delivery/       # (Layer 4) Berkomunikasi dengan luar (HTTP/REST API Handler Fiber)

```

**Alur Kerja Clean Architecture:**

1. **Delivery** menerima _request_ HTTP (misal: POST `/routine`), membaca JSON, lalu menyerahkannya ke Usecase.
2. **Usecase** menerima data, menjalankan logika bisnis (contoh: _Jika durasi tugas lebih dari 2 jam, tolak request_). Jika lolos, Usecase menyuruh Repository untuk menyimpan.
3. **Repository** menerima perintah dari Usecase, menyusun _query_ SQL, dan menyimpannya ke _database_.

### Strategi Eksekusi Portofolio

Sebagai seorang _engineer_ yang bersiap menembus industri, mengadopsi **Clean Architecture** pada proyek karya akhir Anda akan menjadi nilai jual yang sangat masif di mata _recruiter_ atau _senior engineer_ saat wawancara kerja. Menunjukkan pemahaman bahwa kode HTTP tidak boleh dicampuradukkan dengan logika _database_ membuktikan bahwa Anda siap bekerja dalam sistem yang besar dan kolaboratif.

Apakah struktur dan penamaan lapisan-lapisan ini sudah cukup tergambar logikanya? Jika sudah, kita sudah siap untuk masuk ke tahap pamungkas dari _roadmap_ ini, yaitu **Sesi 10: Integrasi Total - Mini Project Portofolio**!
