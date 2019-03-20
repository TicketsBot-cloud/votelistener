package endpoints

import (
	"encoding/json"
	"github.com/TicketsBot/VoteListener/config"
	"github.com/TicketsBot/VoteListener/database"
	"github.com/valyala/fasthttp"
	"strconv"
)

type(
	Request struct {
		Bot string // Snowflake
		User string // Snowflake
		Type string // upvote || test
		IsWeekend bool
	}
)

func VoteHandler(ctx *fasthttp.RequestCtx) {
	auth := ctx.Request.Header.Peek("Authorization")
	if auth == nil {
		ctx.Error("No token specified", 400)
		return
	} else if string(auth) != config.Conf.Bot.Token {
		ctx.Error("Invalid token", 400)
		return
	}

	body := ctx.PostBody()
	if auth == nil {
		ctx.Error("No post data", 400)
		return
	}

	var req Request
	err := json.Unmarshal(body, &req); if err != nil {
		ctx.Error("Invalid request", 400)
		return
	}

	id, err := strconv.ParseInt(req.User, 10, 64); if err != nil {
		ctx.Error("Invalid user ID", 400)
		return
	}

	database.AddVote(id)
}
