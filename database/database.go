package database

import (
	"fmt"
	"github.com/TicketsBot/VoteListener/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type(
	Votes struct {
		Id int64 `gorm:"type:bigint;unique_index;primary_key"`
		VoteTime time.Time
	}
)

var(
	database gorm.DB
)

func ConnectDatabase() {
	log.Println("Connecting to DB")

	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Database.Username,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Database,
	)

	db, err := gorm.Open("mysql", uri); if err != nil {
		panic(err)
	}

	database = *db

	log.Println("Connected to DB")
}

func CreateTables() {
	log.Println("Creating tables")
	database.Exec("CREATE TABLE IF NOT EXISTS votes(id BIGINT UNIQUE PRIMARY KEY, vote_time TIMESTAMP);")
}

func AddVote(userId int64) {
	vote := Votes{
		Id: userId,
		VoteTime: time.Now(),
	}

	database.Where(Votes{Id: userId}).Assign(Votes{VoteTime: time.Now()}).FirstOrCreate(&vote)
}