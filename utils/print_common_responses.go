package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success Data Response

type successResponseForApp struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ShowAppCommonSuccessResponse generates a success JSON response.
func ShowAppCommonSuccessResponse(c *gin.Context, message string, data interface{}) {
	response := &successResponseForApp{
		Success: true,
		Data:    data,
		Message: message,
	}

	c.JSON(http.StatusOK, response)
	return
}

// Error Data Response

type errorResponseForApp struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ShowAppCommonErrorResponse(c *gin.Context, statusCode int, errorMessage string, data interface{}) {
	response := errorResponseForApp{
		Success: false,
		Message: errorMessage,
		Data:    data,
	}

	c.JSON(statusCode, response)
	return
}
