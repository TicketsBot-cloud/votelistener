package endpoints

import  "github.com/valyala/fasthttp"

func Index(ctx *fashttp.RequestCtx) {
	ctx.Redirect("https://top.gg/bot/508391840525975553/vote", 301)
}
