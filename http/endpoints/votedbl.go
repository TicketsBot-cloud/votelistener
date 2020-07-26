package endpoints

import (
	"github.com/TicketsBot/VoteListener/database"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type DBLRequest struct {
	Admin    bool   `json:"admin"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Id       uint64 `json:"id,string"`
}

func DBLHandler(ctx *gin.Context) {
	sig := ctx.GetHeader("X-DBL-Signature")
	split := strings.Split(sig, " ")
	if len(split) != 2 {
		ctx.String(403, "Invalid signature")
		return
	}

	secret := split[0]

	if secret != os.Getenv("DBL_TOKEN") {
		ctx.String(403, "Invalid token")
		return
	}

	var body DBLRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.String(400, err.Error())
		return
	}

	database.AddVote(body.Id)
}
