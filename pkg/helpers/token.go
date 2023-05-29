package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTokenFromRequest(ctx *gin.Context) (string, error) { 
	header := ctx.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}