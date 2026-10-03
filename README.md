# Job Application Tracker

Aplikasi untuk melacak lamaran kerja.

## Fitur

- **Autentikasi**: register, login (JWT), reset password
- **Manajemen lamaran**: tambah, lihat, hapus lamaran kerja
- **Status pipeline**: `wishlist → applied → interview → offer → accepted`, dengan opsi `rejected` di tiap tahap
- **Dashboard**: ringkasan jumlah lamaran per status

## Tech Stack

**Backend**
- Go, Gin (HTTP framework), GORM (ORM)
- PostgreSQL, golang-migrate (migrasi database)
- JWT (golang-jwt), bcrypt (hash password)
- gomail (kirim email), Mailhog

**Frontend**
- Next.js (App Router), TypeScript
- Tailwind CSS
- React Context (state management)


## Menjalankan Secara Lokal

### Prasyarat
- Go 1.23+
- Node.js 18+
- Docker Desktop

### 1. Clone repo

```bash
git clone https://github.com/LouisJonathan88/job-tracker.git
cd job-tracker
```

### 2. Siapkan environment variable

Buat file `.env` di root proyek (lihat `.env.example` untuk daftar variabel):

```bash
cp .env.example .env
```

Isi `JWT_SECRET` dengan string acak yang panjang.

### 3. Jalankan PostgreSQL dan Mailhog

```bash
docker compose up -d
```

### 4. Jalankan migrasi database

```bash
cd backend
migrate -path migrations -database "postgres://<DB_USER>:<DB_PASSWORD>@localhost:<DB_PORT>/<DB_NAME>?sslmode=disable" up
```

### 5. Jalankan backend

```bash
go run ./cmd/api
```

Backend berjalan di `http://localhost:8080`.

### 6. Jalankan frontend

```bash
cd frontend
cp .env.example .env
npm install
npm run dev
```

Frontend berjalan di `http://localhost:3000`.

### 7. Cek email reset password

Buka `http://localhost:8025` untuk melihat email yang terkirim lewat Mailhog.