package controllers

import (
	"github.com/bantawa04/bgo-api/app/response"
	"github.com/bantawa04/bgo-api/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserApiController struct {
	logger config.Logger
	env    config.Env
}

func NewUserController(logger config.Logger, env config.Env) UserApiController {
	return UserApiController{
		logger: logger,
		env:    env,
	}
}

func (cc UserApiController) GetUserProfile(c *gin.Context) {

	userID := c.Param("id")

	// Convert string ID to int
	id, err := strconv.Atoi(userID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// For demonstration, using hardcoded user data
	user := map[string]interface{}{
		"id":    id,
		"name":  "John Doe",
		"email": "john@example.com",
	}

	response.Success(c, "User retrieved successfully", user)
}
