# Rate Limit Testing

Repository untuk testing rate limiting dengan dua komponen: server dan client tester.

## Struktur Project
```
rate-limit-testing/
├── rate-limit-spike-test/     # Server dengan rate limiting
└── test-hit-rate-limit/        # Client untuk test hit server
```

## Cara Menggunakan

### 1. Jalankan Server
```bash
cd rate-limit-spike-test
go run main.go
```

### 2. Jalankan Test Client
```bash
cd test-hit-rate-limit
go run main.go
```

## Deskripsi

- **rate-limit-spike-test**: Server yang mengimplementasikan rate limiting
- **test-hit-rate-limit**: Program untuk melakukan spike test terhadap server dengan multiple requests

## Catatan
Pastikan server sudah berjalan sebelum menjalankan test client.
