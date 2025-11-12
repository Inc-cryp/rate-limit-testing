# Rate Limit Spike Test - Server

Server HTTP sederhana dengan implementasi rate limiting untuk testing.

## Deskripsi

Server ini menyediakan endpoint API yang dapat di-test untuk rate limiting. Digunakan bersama dengan client tester untuk mensimulasikan spike traffic.

## Cara Menjalankan

### Prerequisites
- Go 1.20 atau lebih tinggi

### Jalankan Server
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Endpoint

### GET `/api/books?subject=romance`

Endpoint untuk testing rate limiting.

**Response Success (200 OK):**
```json
{
  "message": "Request berhasil",
  "data": [...]
}
```

**Response Rate Limit (429 Too Many Requests):**
```json
{
  "error": "Rate limit exceeded"
}
```

## Testing

Test dengan curl:
```bash
curl "http://localhost:8080/api/books?subject=romance"
```

Atau gunakan client tester di folder `test-hit-rate-limit` untuk spike testing dengan 50 concurrent requests.

## Konfigurasi Rate Limit

- Rate limit dapat dikonfigurasi sesuai kebutuhan testing
- Default: [sesuaikan dengan implementasi Anda]

## Notes

Server ini dibuat untuk testing purposes.
