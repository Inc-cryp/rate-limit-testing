package books

import (
	"context"
	"encoding/json"
	"log"
	dto "rate-limit-spike-test/src/app/dto/books"
	"time"

	integ "rate-limit-spike-test/src/infra/integration/books"
	redis "rate-limit-spike-test/src/infra/persistence/redis/service"
)

type BooksUCInterface interface {
	GetBooksBySubject(ctx context.Context, req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
}

type booksUseCase struct {
	BooksInteg integ.OpenLibraryServices
	Redis      redis.ServRedisInt
}

func NewBooksUseCase(i integ.OpenLibraryServices, r redis.ServRedisInt) *booksUseCase {
	return &booksUseCase{
		BooksInteg: i,
		Redis:      r,
	}
}

func (uc *booksUseCase) GetBooksBySubject(ctx context.Context, req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error) {

	var resp *dto.GetBooksRespDTO

	dataRedis, err := uc.Redis.GetData(ctx, req.Subject)

	if err != nil {
		log.Printf("unable to GET data from redis. error: %v", err)
	}

	if dataRedis != "" {
		// ambil data dari redis jika ada
		//kenapa pointer "&resp" karena akan di unmarshal ke struct resp artinya ditulis langsung ke alamat memori resp
		_ = json.Unmarshal([]byte(dataRedis), &resp)

		log.Println("data from redis")
		return resp, nil

	}

	resp, err = uc.BooksInteg.GetBooksBySubject(req.Subject)
	log.Println("data not from redis")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	redisData, _ := json.Marshal(resp)
	ttl := time.Duration(2) * time.Minute

	// set data ke redis
	err = uc.Redis.SetData(context.Background(), req.Subject, redisData, ttl)
	if err != nil {
		log.Printf("unable to SET data. error: %v", err)
		return nil, err
	}

	return resp, nil
}
