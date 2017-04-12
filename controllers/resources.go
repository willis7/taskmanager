package controllers

import "github.com/willis7/taskmanager/models"

type (
	// For Post - /users/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	// For Post - /users/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	// Response for authorization user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// Model for authentication
	LoginModel struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	// Model for authorixed user with access token
	AuthUserModel struct {
		User models.User `json:"user"`
		Token string `json:"token"`
	}
)
