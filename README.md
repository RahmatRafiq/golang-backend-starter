
# Golang Backend Starter

## Deskripsi
Golang Backend Starter adalah starter kit untuk memulai pengembangan aplikasi backend menggunakan bahasa Go. Proyek ini menyediakan struktur modular dengan komponen-komponen yang siap digunakan seperti:

- **Controllers**: Mengelola logika bisnis seperti autentikasi, kategori, produk, dan izin.
- **Helpers**: Fungsi bantuan untuk pengelolaan lingkungan, hashing, dan respons.
- **Middleware**: Pengelolaan autentikasi dan logging.
- **Models**: Definisi skema data (pengguna, produk, kategori).
- **Services**: Logika bisnis aplikasi.
- **Routes**: Pengaturan rute aplikasi.

Proyek ini juga terintegrasi dengan Swagger untuk dokumentasi API. Kamu bisa mengakses dokumentasi Swagger melalui Swagger UI setelah menjalankan proyek.

## Fitur Utama
- Autentikasi menggunakan JWT
- Manajemen pengguna dan otoritas
- Hot-reload dengan Air
- Dokumentasi API otomatis dengan Swagger
- Arsitektur modular dengan layanan terpisah
  
## Konfigurasi Environment
Proyek ini menggunakan file .env untuk mengatur konfigurasi, berikut adalah contoh file .env:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret_key
```

## Cara Menjalankan Aplikasi

1. **Clone repository**:
   ```
   git clone https://github.com/RahmatRafiq/golang-backend-starter.git
   ```
2. **Masuk ke direktori project**:
   ```
   cd golang-backend-starter/backend
   ```
3. **Instal dependencies**:
   ```
   go install github.com/cosmtrek/air@latest
   ```
4. **Jalankan aplikasi dengan Air**:
   ```
   air
   ```
5. **Akses Halaman API dengan url**:
   ```
   http://localhost:8080/swagger/index.html
   ``` 

## Kontribusi

Aplikasi ini dikembangkan oleh [Dzyfhuba](https://github.com/Dzyfhuba) dan [RahmatRafiq](https://github.com/RahmatRafiq). Jangan ragu untuk mengajukan pertanyaan atau memberikan saran. Kami sangat terbuka terhadap kontribusi dari siapa saja yang ingin terlibat.

Selamat mencoba, dan semoga proyek ini membantu kamu dalam pengembangan aplikasi backend! 😊
