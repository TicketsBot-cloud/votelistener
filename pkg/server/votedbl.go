package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type DBLRequest struct {
	Admin    bool   `json:"admin"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Id       uint64 `json:"id,string"`
}

func (s *Server) DblHandler(ctx *gin.Context) {
	if len(s.config.DblToken) == 0 || ctx.GetHeader("Authorization") != s.config.DblToken {
		ctx.String(403, "Invalid token")
		return
	}

	var body DBLRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid body",
		})
		return
	}

	if err := s.db.WithTx(ctx, func(tx pgx.Tx) error {
		return s.db.VoteCredits.Increment(ctx, tx, body.Id)
	}); err != nil {
		ctx.JSON(500, gin.H{
			"message": "An error occurred while processing the request",
		})
		return
	}

	s.logger.Info("Received and stored vote", zap.Uint64("user", body.Id))

	ctx.Status(204)
}
