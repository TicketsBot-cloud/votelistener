package http

import (
	"github.com/TicketsBot/VoteListener/config"
	"github.com/TicketsBot/VoteListener/http/endpoints"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func StartServer() {
	router := routing.New()

	voteHandler := fasthttp.CompressHandler(endpoints.VoteHandler)
	router.Post("/vote", func(ctx *routing.Context) error {
		voteHandler(ctx.RequestCtx)
		return nil
	})

	log.Println("Starting server...")
	err := fasthttp.ListenAndServe(config.Conf.Server.Bind, router.HandleRequest); if err != nil {
		panic(err)
	}
}
