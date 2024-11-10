package main

import (
	"github.com/bantawa04/bgo-api/bootstrap"
	"go.uber.org/fx"
)

// func main() {
//
//		router := gin.Default()
//
//		// Example usage with different scenarios
//		router.GET("/", func(c *gin.Context) {
//			response.Success(c, "Hello, World!", nil)
//		})
//
//		router.GET("/user", func(c *gin.Context) {
//			user := map[string]interface{}{
//				"id":    1,
//				"name":  "John Doe",
//				"email": "john@example.com",
//			}
//			response.Success(c, "User retrieved successfully", user)
//		})
//
//		router.GET("/error-example", func(c *gin.Context) {
//			response.Error(c, http.StatusBadRequest, "Something went wrong")
//		})
//
//		router.Run(":8080")
//	}
func main() {
	fx.New(bootstrap.Module).Run()
}
