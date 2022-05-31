package handler

import (
	"encoding/json"
	"fmt"
	"github.com/connormakh/pwHashApi/app/model/response"
	"github.com/connormakh/pwHashApi/app/utils"
	"log"
	"net/http"
)

type StatsHandlerContext struct {
	DbSession *utils.Datastore
}

func (ctx *StatsHandlerContext) GetStats(w http.ResponseWriter, r *http.Request) {
	// Double check it's a post request being made
	r.ParseForm()
	log.Println(r.Form)
	fmt.Println(r.Form.Get("hash"))
	hashResponse := response.GetStatsResponse{
		AverageTime: ctx.DbSession.GetAverageTimeCount(),
		Total:       ctx.DbSession.GetTotal(),
	}
	jsonResponse, err := json.Marshal(hashResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(jsonResponse))

}
