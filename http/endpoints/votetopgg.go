package endpoints

import (
	"github.com/TicketsBot/VoteListener/database"
	"github.com/gin-gonic/gin"
	"os"
)

type TopGGRequest struct {
	Bot       uint64 `json:"bot,string"`
	User      uint64 `json:"user,string"`
	Type      string `json:"type"`
	IsWeekend bool   `json:"isWeekend"`
	Query     string `json:"string"`
}

func TopGGHandler(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth != os.Getenv("TOPGG_TOKEN") {
		ctx.String(403, "Invalid token")
		return
	}

	var body TopGGRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.String(400, err.Error())
		return
	}

	database.AddVote(body.User)
}
