package main

import (
	"context"
	"fase-4-hf-order/external/db/rds/postgres"
	l "fase-4-hf-order/external/logger"
	repositories "fase-4-hf-order/internal/adapters/driven/repositories/sql"
	"fase-4-hf-order/internal/core/application"
	"fase-4-hf-order/internal/core/useCase"
	grpcH "fase-4-hf-order/internal/handler/rpc"
	cp "fase-4-hf-order/order_proto"
	"log"
	"net"
	"os"
	"time"

	"github.com/marcos-dev88/genv"
	"google.golang.org/grpc"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	listener, err := net.Listen("tcp", ":"+os.Getenv("API_RPC_PORT"))

	if err != nil {
		l.Errorf("error to create connection %v", " | ", err)
	}

	defer listener.Close()

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

	h := grpcH.NewHandler(app)

	grpcServer := grpc.NewServer()

	cp.RegisterOrderServer(grpcServer, h.Handler())

	if err := grpcServer.Serve(listener); err != nil {
		l.Errorf("error in server: %v", " | ", err)
	}
}
