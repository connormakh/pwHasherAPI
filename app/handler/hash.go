package handler

import (
	"fmt"
	"github.com/connormakh/pwHashApi/app/model/response"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type HashHandlerContext struct {
	DbSession *utils.Datastore
	Wg        *sync.WaitGroup
}

func (ctx *HashHandlerContext) PostHash(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	err := r.ParseForm()
	if err != nil {
		return
	}
	inputPassword := r.Form.Get("password")
	currentId := ctx.DbSession.Increment()
	fmt.Fprint(w, utils.ToJson(response.NewPostHashResponse(currentId)))
	// run on another thread to avoid delaying response
	ctx.Wg.Add(1)
	go func() {
		defer ctx.Wg.Done()
		time.Sleep(5 * time.Second)
		ctx.DbSession.Insert(currentId, utils.Hash(inputPassword))
		elapsed := time.Since(start)
		ctx.DbSession.CountTime(elapsed.Milliseconds())
	}()

}

func (ctx *HashHandlerContext) GetHash(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/hash/")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, utils.ToJson(response.NewGetHashMalformedIdError()))
		return
	}
	result := ctx.DbSession.FindOne(idInt)
	if result == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, utils.ToJson(response.NewGetHashNotFoundError()))
		return
	}
	hashResponse := response.GetHashResponse{
		PasswordHash: result,
	}
	fmt.Fprint(w, utils.ToJson(hashResponse))
}
