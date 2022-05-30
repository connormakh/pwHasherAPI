package app

import (
	"github.com/connormakh/pwHashApi/app/router"
	"github.com/connormakh/pwHashApi/app/utils"
)

// App has router and db instances
type App struct {
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	db = utils.Datastore{}.Initialize()
	router.SetupHttpListeners()
}