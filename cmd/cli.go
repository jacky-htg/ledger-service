package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"ledger-service/internal/config"
	"ledger-service/internal/pkg/db/postgres"
	"ledger-service/internal/scheme"
)

func main() {
	if _, ok := os.LookupEnv("PORT"); !ok {
		config.Setup(".env")
	}

	// =========================================================================
	// Logging
	log := log.New(os.Stdout, "LMS : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		log.Fatalf("error: shutting down: %s", err)
	}
}

func run(log *log.Logger) error {
	// =========================================================================
	// App Starting

	log.Printf("main : Started")
	defer log.Println("main : Completed")

	// =========================================================================

	// Start Database

	db, err := postgres.Open()
	if err != nil {
		return fmt.Errorf("connecting to db: %v", err)
	}
	defer db.Close()

	// Handle cli command
	flag.Parse()

	switch flag.Arg(0) {
	case "migrate":
		if err := scheme.Migrate(db); err != nil {
			return fmt.Errorf("applying migrations: %v", err)
		}
		log.Println("Migrations complete")
		return nil

	case "seed":
		if err := scheme.Seed(db); err != nil {
			return fmt.Errorf("seeding database: %v", err)
		}
		log.Println("Seed data complete")
		return nil
	}

	return nil
}
