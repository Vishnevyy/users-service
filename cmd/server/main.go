package main

import (
	"log"

	"github.com/Vishnevyy/users-service/internal/database"
	transportgrpc "github.com/Vishnevyy/users-service/internal/transport/grpc"
	"github.com/Vishnevyy/users-service/internal/user"
)

func main() {
	database.InitDB()

	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
