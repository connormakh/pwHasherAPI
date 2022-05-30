package router
import (
	"github.com/connormakh/pwHashApi/app/handler"
	"github.com/connormakh/pwHashApi/app/utils"
	"net/http"
)
const PostUrl = "/hash"

func HandleHashRequests() {
	utils.Post(PostUrl, http.HandlerFunc(handler.PostHash))
}
