package router
import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)
const GetUrl = "/stats"

func HandleStatsRequests() {
	utils.Get(GetUrl, http.HandlerFunc(handler.GetStats))
}
