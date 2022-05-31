package router

import (
	"github.com/connormakh/pwHashApi/app/utils"
	"log"
	"net/http"
)

func SetupHttpListeners(db *utils.Datastore) {
	HandleHashRequests(db)
	HandleStatsRequests(db)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

