package database

import (
	"context"
	"github.com/TicketsBot/VoteListener/config"
	"github.com/TicketsBot/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"time"
)

type(
	Votes struct {
		Id int64 `gorm:"type:bigint;unique_index;primary_key"`
		VoteTime time.Time
	}
)

var(
	client *database.Database
)

func ConnectDatabase() {
	log.Println("Connecting to DB")

	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err != nil {
		panic(err)
	}

	client = database.NewDatabase(pool)

	log.Println("Connected to DB")
}

func AddVote(userId uint64) {
	if err := client.Votes.Set(userId); err != nil {
		log.Println(err.Error())
	}
}