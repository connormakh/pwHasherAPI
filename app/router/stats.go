package router
import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)
const getUrl = "/stats"


func HandleStatsRequests(db *utils.Datastore) {
	h := handler.StatsHandlerContext{DbSession: db}

	utils.Get(getUrl, http.HandlerFunc(h.GetStats), db)
}
