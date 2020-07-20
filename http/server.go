package http

import (
	"github.com/TicketsBot/VoteListener/config"
	"github.com/TicketsBot/VoteListener/http/endpoints"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	router := gin.Default()

	router.GET("/", endpoints.IndexHandler)
	//router.POST("/vote/topgg", endpoints.TopGGHandler)
	router.POST("/vote/dbl", endpoints.DBLHandler)

	log.Println("Starting server...")

	if err := router.Run(config.Conf.Server.Bind); err != nil {
		panic(err)
	}
}
