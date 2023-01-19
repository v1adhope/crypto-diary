package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

// TODO
type userRoutes struct {
	l *logger.Logger
}

func newUserRoutes(handler *gin.RouterGroup, l *logger.Logger) {

	r := &positionRoutes{
		l: l,
	}
	_ = r

	h := handler.Group("/auth")
	{
		h.POST("/sign-up")
		h.POST("/sign-in")
	}
}
