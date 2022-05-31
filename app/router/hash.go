package router
import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)
const PostUrl = "/hash"
const GetUrl = "/hash/"

func HandleHashRequests(db *utils.Datastore) {
	h := handler.HashHandlerContext{DbSession: db}

	utils.Post(PostUrl, http.HandlerFunc(h.PostHash), db)
	utils.Get(GetUrl, http.HandlerFunc(h.GetHash), db)
}