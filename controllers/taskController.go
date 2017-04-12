package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/willis7/taskmanager/common"
	"github.com/willis7/taskmanager/data"
)

// Handler for HTTP Post - /tasks
// Insert a new Task document
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			http.StatusInternalServerError,
		)
		return
	}
	task := &dataResource
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	// Insert a Task document
	repo.Create(task)
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected Error has occurred",
			http.StatusInternalServerError,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// Handler for HTTP Get - /tasks
// Returns all task documents
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c:= context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	tasks := repo.GetAll()
	if j, err := json.Marshal(TaskResource{Data: tasks}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected Error has occurred",
			http.StatusInternalServerError,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
