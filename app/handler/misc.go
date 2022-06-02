package handler

import (
	"net/http"
	"os"
)

type MiscHandlerContext struct {
	Channel chan os.Signal
}

func (ctx *MiscHandlerContext) PostShutdown(w http.ResponseWriter, r *http.Request) {
	ctx.Channel <- os.Interrupt
}
