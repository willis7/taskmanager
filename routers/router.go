package routers

import "github.com/gorilla/mux"

// InitRoutes initializes the routes for all the project entities
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User, Task and Note entities
	router = SetUserRoutes(router)
	router = SetTaskRoutes(router)
	router = SetNoteRoutes(router)
	return router
}
