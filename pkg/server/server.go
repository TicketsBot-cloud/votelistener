package server

import (
	"os"

	"github.com/TicketsBot-cloud/VoteListener/pkg/config"
	"github.com/TicketsBot-cloud/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		ctx.Redirect(302, "https://ticketsbot.cloud/vote")
	})

	router.POST("/vote/dbl", s.DblHandler)
	router.POST("/vote/tgg", s.TggHandler)

	s.logger.Info("Starting server")

	if err := router.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}
