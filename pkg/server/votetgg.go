package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type TGGRequest struct {
	User uint64 `json:"id,string"`
}

func (s *Server) TggHandler(ctx *gin.Context) {
	if len(s.config.DblToken) == 0 || ctx.GetHeader("Authorization") != s.config.TggToken {
		ctx.String(403, "Invalid token")
		return
	}

	var body TGGRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid body",
		})
		return
	}

	if err := s.db.WithTx(ctx, func(tx pgx.Tx) error {
		return s.db.VoteCredits.Increment(ctx, tx, body.User)
	}); err != nil {
		ctx.JSON(500, gin.H{
			"message": "An error occurred while processing the request",
		})
		return
	}

	s.logger.Info("Received and stored vote", zap.Uint64("user", body.User))

	ctx.Status(204)
}
