<h1 align="center">🏥 Sentra Medika - Backend API</h1>

**Sentra Medika** adalah sistem backend RESTful API untuk manajemen rekam medis elektronik. Sistem ini dirancang untuk mendigitalkan proses pencatatan kesehatan dengan keamanan data yang ketat menggunakan otentikasi JWT dan pembagian hak akses (Role-Based Access Control).

Proyek ini dibangun sebagai implementasi dari kebutuhan sistem informasi klinik modern yang memisahkan peran antara Administrator, Dokter, dan Pasien.

## 🚀 Tech Stack

Aplikasi ini dibangun menggunakan teknologi modern dan _best practice_ dalam ekosistem Golang:

- **Language:** Go (Golang)
- **Framework:** Gin Gonic (High-performance HTTP Web Framework)
- **Database:** PostgreSQL
- **ORM:** GORM (Object Relational Mapping)
- **Authentication:** JWT (JSON Web Token) + Bcrypt Hashing
- **ID Management:** Google UUID (v4)
- **Testing:** Go Test + Testify
- **Architecture:** Modular Clean Architecture (Services, Models, Utils, Handlers)

## 📋 Fitur Utama (Berdasarkan SRS)

Sistem ini memiliki 3 aktor utama dengan hak akses yang berbeda:

### 1. 🔐 Authentication & Security

- **Register:** Pendaftaran akun baru (Hash password otomatis dengan Bcrypt).
- **Login:** Otentikasi aman menggunakan JWT Access Token (15 menit) & Refresh Token (7 hari).
- **Token Rotation:** Fitur refresh token untuk memperbarui sesi tanpa login ulang.
- **Middleware:** Proteksi endpoint berdasarkan validitas token dan Role Guard.

### 2. 👤 Administrator (Admin)

- Memiliki akses penuh untuk manajemen pengguna.
- **CRUD Users:** Melihat, membuat, mengedit, dan menghapus data user (Admin, Dokter, Pasien).
- **Validasi Data:** Memastikan integritas data (email unik, format valid).

### 3. 👨‍⚕️ Dokter (Doctor)

- Memiliki wewenang medis untuk mengelola rekam medis pasien.
- **Create Record:** Membuat diagnosa dan rencana pengobatan untuk pasien.
- **Update Record:** Memperbarui catatan medis jika ada perubahan kondisi.
- **View Records:** Melihat riwayat medis pasien untuk keperluan diagnosa.

### 4. 🤒 Pasien (Patient)

- Memiliki akses terbatas hanya untuk data pribadi.
- **View My Records:** Hanya bisa melihat riwayat medis miliknya sendiri (Privacy First).

## 🛠️ Instalasi & Cara Menjalankan

Ikuti langkah berikut untuk menjalankan proyek di lokal komputer Anda.

### Prasyarat

- Go (Versi 1.20+)
- PostgreSQL (sudah terinstall dan berjalan)

### Langkah-langkah

1.  **Clone Repository**

    ```bash
    git clone https://github.com/a6iyyu/go.git
    cd go
    ```

2.  **Setup Database**

    - Buat database baru di PostgreSQL bernama `sentra-medika`.
    - Pastikan extension `uuid-ossp` aktif (biasanya otomatis dihandle oleh aplikasi).

3.  **Konfigurasi Environment (.env)**
    Buat file `.env` di root folder dan isi sesuai konfigurasi database Anda:

    ```env
    DATABASE_URL="postgresql://postgres:passwordmu@localhost:5432/sentra-medika"
    JWT_SECRET="rahasia_super_aman_panjang_banget_stringnya"
    ```

4.  **Install Dependencies**

    ```bash
    go mod tidy
    ```

5.  **Seeding Data Awal (Wajib untuk pertama kali)**
    Isi database dengan data dummy (Admin, Dokter, Pasien) dan data medis awal:

    ```bash
    go run main.go -seed
    ```

6.  **Jalankan Server**
    ```bash
    go run main.go
    ```
    Server akan berjalan di `http://localhost:7000`

## 🧪 Testing

Proyek ini dilengkapi dengan **Integration Testing** untuk memastikan semua modul (Auth, User CRUD) berjalan dengan benar. Test mencakup skenario positif dan negatif.

Jalankan perintah berikut untuk memulai testing:

```bash
# Menjalankan semua test di seluruh folder
go test .\tests\ -v
```

## 📡 API Endpoints Documentation

Berikut adalah daftar endpoint yang tersedia:

### 🔓 Public / Auth

| Method | Endpoint        | Deskripsi                           |
| :----- | :-------------- | :---------------------------------- |
| `POST` | `/auth/login`   | Masuk ke sistem (mendapatkan Token) |
| `POST` | `/auth/refresh` | Mendapatkan Access Token baru       |
| `POST` | `/auth/logout`  | Keluar dari sistem (revoke token)   |

### 🛡️ Admin Only (`Bearer Token` + Role `admin`)

| Method   | Endpoint           | Deskripsi                         |
| :------- | :----------------- | :-------------------------------- |
| `POST`   | `/admin/register`  | Mendaftarkan akun admin baru      |
| `POST`   | `/admin/users`     | Membuat user baru (Dokter/Pasien) |
| `GET`    | `/admin/users`     | Melihat daftar semua user         |
| `PUT`    | `/admin/users/:id` | Mengedit data user                |
| `DELETE` | `/admin/users/:id` | Menghapus user                    |

### 🩺 Doctor Only (`Bearer Token` + Role `doctor`)

| Method   | Endpoint               | Deskripsi                 |
| :------- | :--------------------- | :------------------------ |
| `POST`   | `/medical/records`     | Membuat rekam medis baru  |
| `GET`    | `/medical/records`     | Melihat semua rekam medis |
| `PUT`    | `/medical/records/:id` | Mengupdate rekam medis    |
| `DELETE` | `/medical/records/:id` | Menghapus rekam medis     |

### 👤 Patient Only (`Bearer Token` + Role `patient`)

| Method | Endpoint              | Deskripsi                          |
| :----- | :-------------------- | :--------------------------------- |
| `GET`  | `/medical/my-records` | Melihat riwayat sakit diri sendiri |

## 📂 Struktur Proyek

```bash
sentra-medika/
├── main.go # Entry point aplikasi
├── go.mod # Dependency manager
├── .env # Environment variables
├── models/ # Definisi Struct & Schema Database (GORM)
├── services/ # Business Logic & Handlers
├── seeders/ # Script pengisi data awal
├── utils/ # Helper (Database connect, Hash, JWT)
├── constants/ # Konstanta URL Path
└── tests/ # Integration Tests
```

**Dibuat dengan ❤️**