package main

import (
	"github.com/willis7/taskmanager/common"
	"github.com/willis7/taskmanager/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main() {
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
