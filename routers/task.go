package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/willis7/taskmanager/common"
)

// SetTaskRoutes registers the Task routes
func SetTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.GetTaskById).Methods("GET")
	taskRouter.HandleFunc("/tasks/users/{id}", controllers.GetTaskByUser).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	// applies the middleware handler function named Authorize to all routes that map to URI "/tasks"
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}
