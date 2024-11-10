// response/response.go
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"code"`
}

func Success(c *gin.Context, message string, data interface{}) {
	response := Response{
		Status:     true,
		Message:    message,
		Data:       data,
		StatusCode: http.StatusOK,
	}
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, statusCode int, message string) {
	response := Response{
		Status:     false,
		Message:    message,
		StatusCode: statusCode,
	}
	c.JSON(statusCode, response)
}
