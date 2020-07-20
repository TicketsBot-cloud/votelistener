package endpoints

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.Redirect(302, "https://discordbotlist.com/bots/tickets/upvote")
	/*if rand.Intn(10) <= 5 {
		ctx.Redirect(302, "https://top.gg/bot/508391840525975553/vote")
	} else {
		ctx.Redirect(302, "https://discordbotlist.com/bots/tickets/upvote")
	}*/
}
