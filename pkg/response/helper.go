package response

import (
	"github.com/gin-gonic/gin"
	"poymanov/todo/pkg/helpers"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{helpers.FirstToUpper(message)})
}

type ErrorResponse struct {
	Message string `json:"message"`
}
