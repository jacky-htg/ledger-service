package route

import (
	"database/sql"
	"log"

	"ledger-service/internal/domain/coa"
	"ledger-service/pb/ledgers"

	"google.golang.org/grpc"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, db *sql.DB, log *log.Logger) {
	coaServer := coa.Service{Db: db}
	ledgers.RegisterChartOfAccountServiceServer(grpcServer, &coaServer)
}
