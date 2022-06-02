package router

import (
	"github.com/connormakh/pwHashApi/app/model"
)

func SetupHttpListeners(serverContext *model.ServerContext) {
	HandleHashRequests(serverContext)
	HandleStatsRequests(serverContext)
	HandleMiscRequests(serverContext)

	//log.Fatal(http.ListenAndServe(":10000", nil))
}

