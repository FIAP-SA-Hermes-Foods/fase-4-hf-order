package main

import (
	"context"
	"fase-4-hf-order/external/db/rds/postgres"
	l "fase-4-hf-order/external/logger"
	repositories "fase-4-hf-order/internal/adapters/driven/repositories/sql"
	"fase-4-hf-order/internal/core/application"
	"fase-4-hf-order/internal/core/useCase"
	httpH "fase-4-hf-order/internal/handler/http"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()
	ctx := context.Background()

	dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

	if err != nil {
		log.Fatalf("error: %v", err)

	}

	db := postgres.NewPostgresDB(
		ctx,
		os.Getenv("DB_REGION"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		dbDuration,
	)

	repoOrder := repositories.NewOrderDB(ctx, db)

	repoOrderItem := repositories.NewOrderItemDB(ctx, db)

	uc := useCase.NewOrderUseCase()

	app := application.NewApplication(repoOrder, repoOrderItem, uc)

	h := httpH.NewHandler(app)
	router.Handle("/hermes_foods/health", http.StripPrefix("/", httpH.Middleware(h.HealthCheck)))
	router.Handle("/hermes_foods/order/", http.StripPrefix("/", httpH.Middleware(h.Handler)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
