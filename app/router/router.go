package router

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func SetupHttpListeners() {
	HandleHashRequests()
	HandleStatsRequests()
	log.Fatal(http.ListenAndServe(":10000", nil))
}

