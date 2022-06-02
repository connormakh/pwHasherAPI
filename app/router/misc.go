package router
import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/model"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)
const shutdownUrl = "/shutdown"

func HandleMiscRequests(serverContext *model.ServerContext) {
	h := handler.MiscHandlerContext{Channel: serverContext.Channel}
	utils.Post(shutdownUrl, http.HandlerFunc(h.PostShutdown))
}