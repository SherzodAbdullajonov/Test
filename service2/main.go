package main

import (
	"context"
	"github.com/SherzodAbdullajonov/service2/router"
	"golang_projects/test-crud/common/environment"
	"golang_projects/test-crud/common/postgres"
	"golang_projects/test-crud/common/server"
	"log"
	"time"

	"github.com/SherzodAbdullajonov/service2/application"
	"github.com/SherzodAbdullajonov/service2/pb"

	"github.com/SherzodAbdullajonov/service2/database"
	"github.com/SherzodAbdullajonov/service2/repository"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cfg, err := database.Load()
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}
	
	
	db, err := postgres.Connect(ctx, postgres.Config{ConnString: cfg.PostgresConnString()})
	if err != nil {
		log.Panicf("failed to connect to db: %v; connString: %s", err, cfg.PostgresConnString())
	}
	defer db.Close()

	app := application.New(repository.NewPostgresRepository(db))
	server.RunGRPCServer(cfg.ListenAddress(), func(s *grpc.Server) {
		svc := router.NewGrpcServer(app)
		pb.RegisterPostServiceServer(s, svc)
	})
}
