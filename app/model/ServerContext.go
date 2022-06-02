package model

import (
	"github.com/connormakh/pwHashApi/app/utils"
	"os"
	"sync"
)

type ServerContext struct {
	Db *utils.Datastore
	Channel chan os.Signal
	Wg *sync.WaitGroup
}
