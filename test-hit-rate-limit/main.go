package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// Buat rate limiter dengan batas 50 permintaan per detik
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*20), 1)

	// WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Buat 50 permintaan secara bersamaan
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(id int) {
			// Kurangi WaitGroup ketika permintaan selesai
			defer wg.Done()

			// Tunggu hingga diizinkan oleh rate limiter
			limiter.Wait(context.Background())

			// Tampilkan waktu pengiriman permintaan
			fmt.Printf("Request %d sent at %s\n", id, time.Now().Format("15:04:05.000"))

			// Buat permintaan ke API
			err := makeRequest()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Request succeeded!")
			}

		}(i)
	}

	// Tunggu hingga semua permintaan selesai
	wg.Wait()
}

func makeRequest() error {
	url := "http://localhost:8080/api/books?subject=love"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
