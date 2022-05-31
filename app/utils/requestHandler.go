package utils

import (
	"net/http"
)

func handlerWrapper(path string, handler http.Handler, method string, db *Datastore) {
	handleFunction := func (w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	}
	http.Handle(path, http.HandlerFunc(handleFunction))
}

func Get(path string, handler http.Handler, db *Datastore) {
	handlerWrapper(path, handler, http.MethodGet, db)
}

func Post(path string, handler http.Handler, db *Datastore) {
	handlerWrapper(path, handler, http.MethodPost, db)
}