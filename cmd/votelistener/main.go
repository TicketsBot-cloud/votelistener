package main

import (
	"context"
	"time"

	"github.com/TicketsBot-cloud/VoteListener/pkg/config"
	"github.com/TicketsBot-cloud/VoteListener/pkg/server"
	"github.com/TicketsBot-cloud/common/observability"
	"github.com/TicketsBot-cloud/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}

	logger, err := observability.Configure(cfg.SentryDsn, cfg.JsonLogs, cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	logger.Info("Connecting to database...")
	db := connectToDatabase(cfg)
	logger.Info("Connected to database")

	s := server.NewServer(logger, cfg, db)
	s.Run()
}

func connectToDatabase(cfg config.Config) *database.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, cfg.DatabaseUri)
	if err != nil {
		panic(err)
	}

	return database.NewDatabase(pool)
}
