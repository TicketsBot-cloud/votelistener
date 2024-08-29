package server

import (
	"github.com/TicketsBot/VoteListener/pkg/config"
	"github.com/TicketsBot/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

type Server struct {
	logger *zap.Logger
	config config.Config
	db     *database.Database
}

func NewServer(logger *zap.Logger, config config.Config, db *database.Database) *Server {
	return &Server{
		logger: logger,
		config: config,
		db:     db,
	}
}

func (s *Server) Run() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "https://discordbotlist.com/bots/tickets/upvote")
	})

	router.POST("/vote/dbl", s.DblHandler)

	s.logger.Info("Starting server")

	if err := router.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}
