package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

// TODO
type positionRoutes struct {
	l *logger.Logger
}

func newPositionRoutes(handler *gin.RouterGroup, l *logger.Logger) {

	r := &positionRoutes{
		l: l,
	}
	_ = r

	h := handler.Group("/positon")
	{
		h.GET("/")
		h.GET("/:id")
		h.POST("/")
		h.PUT("/")
		h.DELETE("/")
	}
}
