package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/willis7/taskmanager/common"
	"github.com/willis7/taskmanager/models"
	"github.com/willis7/taskmanager/data"
)

// Handler for HTTP Post - /users/register
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			http.StatusInternalServerError,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	// Insert User document
	repo.CreateUser(user)
	// Clean up the hashpassword to eliminate it from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// Handler for HTTP Post /users/login
// Authenticate with username and password
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	var token string
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			http.StatusInternalServerError,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}

	// Authenticate the login user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	} else { // if login is successful
		// Generate JWT token
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the access token",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User: user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				http.StatusInternalServerError,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
