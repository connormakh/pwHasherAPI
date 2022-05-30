package handler

import (
	"fmt"
	"log"
	"net/http"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	// Double check it's a post request being made
	r.ParseForm()
	log.Println(r.Form)
	fmt.Println(r.Form.Get("hash"))
}
