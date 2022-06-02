package launcher

import (
	"context"
	"github.com/connormakh/pwHashApi/app/model"
	"github.com/connormakh/pwHashApi/app/router"
	"github.com/connormakh/pwHashApi/app/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// App has router and db instances
type App struct {
	ServerContext *model.ServerContext
}

// Initialize initializes the app with predefined configuration
func Initialize(channel chan os.Signal, wg *sync.WaitGroup) App {
	serverContext := model.ServerContext{
		Db:      utils.InitializeDatastore(),
		Channel: make(chan os.Signal, 1),
		Wg:      &sync.WaitGroup{},
	}
	app := App{ServerContext: &serverContext}
	router.SetupHttpListeners(app.ServerContext)

	signal.Notify(channel, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-channel
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx, wg); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

	return app
}

func serve(ctx context.Context, wg *sync.WaitGroup) (err error) {

	srv := &http.Server{
		Addr:    ":10000",
		Handler: nil,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()
	wg.Wait()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
