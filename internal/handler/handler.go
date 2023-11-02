package handler

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Register(r *gin.Engine)
}
