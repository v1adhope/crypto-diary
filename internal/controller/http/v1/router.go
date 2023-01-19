package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

// TODO
func NewRouter(handler *gin.Engine, l *logger.Logger) {
	h := handler.Group("/v1")
	{
		newPositionRoutes(h, l)
		newUserRoutes(h, l)
	}
}
