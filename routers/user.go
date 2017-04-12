package routers

import "github.com/gorilla/mux"

// SetUserRoutes receives a pointer to the Gorilla mux router object (mux.Router), specifies 2 routes
// for Login and Register, then returns the pointer of mux.Router object
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
