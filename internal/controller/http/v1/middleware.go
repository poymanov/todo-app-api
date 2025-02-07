package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"poymanov/todo/pkg/response"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	ContextEmailKey     = "ContextEmailKey"
)

func (h *Handler) auth(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if !strings.HasPrefix(header, "Bearer ") {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	token := strings.TrimPrefix(header, "Bearer ")

	isValid, data := h.jwt.Parse(token)

	if !isValid {
		response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	c.Set(ContextEmailKey, data.Email)
}

func getContextEmail(c *gin.Context) (string, error) {
	contextValue, ok := c.Get(ContextEmailKey)
	if !ok {
		return "", errors.New("email not found")
	}

	email, ok := contextValue.(string)
	if !ok {
		return "", errors.New("email is of invalid type")
	}

	return email, nil
}
