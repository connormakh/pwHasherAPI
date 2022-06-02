package router
import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/model"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)
const getUrl = "/stats"


func HandleStatsRequests(serverContext *model.ServerContext) {
	h := handler.StatsHandlerContext{DbSession: serverContext.Db}

	utils.Get(getUrl, http.HandlerFunc(h.GetStats))
}
