package main

import (
	"github.com/TicketsBot/VoteListener/config"
	"github.com/TicketsBot/VoteListener/database"
	"github.com/TicketsBot/VoteListener/http"
)

func main() {
	config.LoadConfig()
	database.ConnectDatabase()
	database.CreateTables()
	http.StartServer()
}
