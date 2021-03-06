package mpapp

import (
	"github.com/julienschmidt/httprouter"
	"github.com/strongo/bots-framework/core"
	"github.com/prizarena/matching-pennies/server-go/mpbot"
)

func InitApp(httpRouter *httprouter.Router, botHost bots.BotHost) {
	mpbot.InitBot(httpRouter, botHost, TheAppContext)
}
