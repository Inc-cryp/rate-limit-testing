# Go Rate Limiter Demo

Program sederhana untuk mendemonstrasikan penggunaan rate limiter dalam mengirim HTTP request secara konkuren menggunakan Go.

## Deskripsi

Program ini mengirimkan 50 HTTP request secara bersamaan ke sebuah API endpoint, dengan rate limiter yang membatasi maksimal 50 request per detik (1 request setiap 20 milidetik).

## Fitur

- **Rate Limiting**: Membatasi kecepatan pengiriman request untuk menghindari overload server
- **Concurrent Requests**: Menggunakan goroutines untuk mengirim multiple request secara bersamaan
- **Logging**: Menampilkan waktu pengiriman setiap request dengan format timestamp

## Cara Menggunakan

### Prasyarat

- Go 1.16 atau lebih tinggi
- Server API yang berjalan di `http://localhost:8080`

### Instalasi Dependencies

```bash
go get golang.org/x/time/rate
```

### Menjalankan Program

```bash
go run main.go
```

## Konfigurasi

### Rate Limiter

```go
limiter := rate.NewLimiter(rate.Every(time.Millisecond*20), 1)
```

- Rate: 1 request setiap 20 milidetik (50 request/detik)
- Burst: 1 (jumlah request yang bisa dilakukan secara burst)

### Endpoint API

```go
url := "http://localhost:8080/api/books?subject=love"
```

Ubah URL sesuai dengan endpoint API yang ingin Anda test.

### Jumlah Request

```go
for i := 1; i <= 50; i++ {
```

Ubah angka `50` untuk mengubah jumlah total request yang akan dikirim.

## Cara Kerja

1. Program membuat rate limiter dengan batas 50 request per detik
2. Membuat 50 goroutines yang masing-masing akan mengirim 1 HTTP request
3. Setiap goroutine menunggu giliran dari rate limiter sebelum mengirim request
4. Request dikirim ke endpoint API dan hasilnya ditampilkan
5. Program menunggu hingga semua request selesai menggunakan WaitGroup

## Output

```
Request 1 sent at 14:30:25.100
Request succeeded!
Request 2 sent at 14:30:25.120
Request succeeded!
Request 3 sent at 14:30:25.140
Request succeeded!
...
```

## Catatan

- Pastikan server API target sudah berjalan sebelum menjalankan program
- Rate limiter membantu mencegah server kewalahan dengan terlalu banyak request sekaligus
- Program menggunakan `sync.WaitGroup` untuk memastikan semua goroutines selesai sebelum program berakhir