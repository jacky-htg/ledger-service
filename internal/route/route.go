package route

import (
	"database/sql"
	"log"

	"google.golang.org/grpc"

	"ledger-service/internal/pkg/db/redis"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, db *sql.DB, log *log.Logger, cache *redis.Cache) {
	//quizServer := quizDomain.QuizService{Db: db, Cache: cache}
	//quizPb.RegisterQuizzesServer(grpcServer, &quizServer)
}
