package launcher

import (
	"github.com/connormakh/pwHashApi/app/handler"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestInitializeShutdown(t *testing.T) {
	c := make(chan os.Signal, 1)
	wg := &sync.WaitGroup{}
	go Initialize(c, wg) // begin initialization
	time.Sleep(3 * time.Second)
	req, _ := http.NewRequest(http.MethodPost, "/shutdown", nil)
	h := handler.MiscHandlerContext{Channel: c}
	h.PostShutdown(nil, req)
}
