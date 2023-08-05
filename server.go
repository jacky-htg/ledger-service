package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"ledger-service/internal/config"
	"ledger-service/internal/middleware"
	"ledger-service/internal/pkg/db/postgres"
	"ledger-service/internal/pkg/db/redis"
	"ledger-service/internal/route"
)

const defaultPort = "8000"

func main() {
	// lookup and setup env
	if _, ok := os.LookupEnv("PORT"); !ok {
		config.Setup(".env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// init log
	log := log.New(os.Stdout, "LMS : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// create postgres database connection
	db, err := postgres.Open()
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
		return
	}
	log.Print("connecting to postgresql database")

	defer db.Close()

	// create redis cache connection
	cache, err := redis.NewCache(context.Background(), os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"), 24*time.Hour)
	if err != nil {
		log.Fatalf("cannot create redis connection: %v", err)
		return
	}
	log.Print("connecting to redis cache")

	// listen tcp port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := middleware.Context{}
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			authInterceptor.Unary(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			authInterceptor.Stream(),
		)),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	// routing grpc services
	route.GrpcRoute(grpcServer, db, log, cache)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	log.Print("serve grpc on port: " + port)
}
