package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petProject/internal/app/errors"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (r *Repository) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, errors.EmptyAuthHeader)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, errors.InvalidAuthHeader)
	}

	userId, err := r.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}
