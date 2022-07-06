package main

import (
	"context"
	"github.com/SherzodAbdullajonov/service1/database"
	"log"
	"time"

	"github.com/SherzodAbdullajonov/service1/application"
	"github.com/SherzodAbdullajonov/service1/pb"
	"github.com/SherzodAbdullajonov/service1/repository"
	"github.com/SherzodAbdullajonov/service1/router"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)
type Config struct {
	ConnString string
}

func Connect(ctx context.Context, cfg Config) (*sqlx.DB, error) {
	return sqlx.ConnectContext(ctx, "postgres", cfg.ConnString)
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cfg := database.Load()
	db, err := Connect(ctx, Config{ConnString: cfg.PostgresConnString()})
	if err != nil {
		log.Panicf("failed to connect to db: %v; connString: %s", err, cfg.PostgresConnString())
	}
	defer db.Close()

	app := application.New(repository.NewPostgresRepository(db))
	server.RunGRPCServer(cfg.ListenAddress(), func(s *grpc.Server) {
		svc := repository.NewGrpcServer(app)
		pb.RegisterDataServiceServer(s, svc)
	})
}
