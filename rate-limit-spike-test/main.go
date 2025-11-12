package main

import (
	"context"

	usecases "rate-limit-spike-test/src/app/usecase"
	bookUC "rate-limit-spike-test/src/app/usecase/books"
	"rate-limit-spike-test/src/infra/config"
	"rate-limit-spike-test/src/infra/persistence/redis"

	"rate-limit-spike-test/src/interface/rest"

	bookInteg "rate-limit-spike-test/src/infra/integration/books"

	ms_log "rate-limit-spike-test/src/infra/log"

	circuit_breaker_service "rate-limit-spike-test/src/infra/circuit_breaker"
	redisService "rate-limit-spike-test/src/infra/persistence/redis/service"

	// "rate-limit-spike-test/src/infra/broker/nats"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// init context ini untuk apa? yaitu untuk menghandle lifecycle aplikasi ini
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// cek production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup ini untuk logging aplikasi
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	redisClient, _ := redis.NewRedisClient(conf.Redis, logger)
	redisSvc := redisService.NewServRedis(redisClient)

	circuitBreaker := circuit_breaker_service.NewCircuitBreakerInstance()
	bookIntegration := bookInteg.NewIntegOpenLibrary(circuitBreaker)

	// NATS Publisher Setup, ini untuk mengirim pesan ke NATS
	// tugas NATS Publisher ini untuk mengirim pesan ke service lain secara asynchronous
	// cara bekerjanya ? yaitu dengan publish-subscribe pattern
	// contohnya:
	// misal ada service pickup yang mengirim pesan ke service notification
	// nah service notification ini yang akan mendengarkan pesan dari service pickup
	// jadi ketika ada event tertentu di service pickup, misal user sudah membuat pickup
	// maka service pickup akan mengirim pesan ke NATS
	// dan service notification akan mendengarkan pesan tersebut dan mengirim notifikasi ke user
	// ini sangat berguna untuk mengurangi beban service utama dan membuat aplikasi lebih scalable
	// Nats := nats.NewNats(conf.Nats, logger)
	// publisher := natsPublisher.NewPushWorker(Nats)

	// HTTP Handler
	// server siap mengimplementasikan semua use case graceful shutdown
	allUC := usecases.AllUseCases{
		BookUC: bookUC.NewBooksUseCase(bookIntegration, redisSvc),
		// PickUpUC: pickUpUC.NewPickUpUseCase(publisher),
	}

	//ini server http untuk menjalankan aplikasi
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		allUC,
		conf.RPS,
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
