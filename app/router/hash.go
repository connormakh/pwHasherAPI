package router

import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/model"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)

const PostUrl = "/hash"
const GetUrl = "/hash/"

func HandleHashRequests(serverContext *model.ServerContext) {
	h := handler.HashHandlerContext{DbSession: serverContext.Db, Wg: serverContext.Wg}

	utils.Post(PostUrl, http.HandlerFunc(h.PostHash))
	utils.Get(GetUrl, http.HandlerFunc(h.GetHash))
}
