package routes

import (
	"github.com/bantawa04/bgo-api/app/controllers"
	"github.com/bantawa04/bgo-api/config"
)

// UserRoutes struct
type UserRoutes struct {
	logger         config.Logger
	router         config.Router
	userController controllers.UserApiController
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger config.Logger,
	router config.Router,
	userController controllers.UserApiController,
) UserRoutes {
	return UserRoutes{
		router:         router,
		logger:         logger,
		userController: userController,
	}
}

// Setup user routes
func (i UserRoutes) Setup() {
	i.logger.Zap.Info(" Setting up user routes")

	i.router.Gin.GET("/profile", i.userController.GetUserProfile)
}
