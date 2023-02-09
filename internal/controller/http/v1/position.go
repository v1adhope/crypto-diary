package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

// TODO
type positionRoutes struct {
	l *logger.Log
}

func newPositionRoutes(handler *gin.RouterGroup, u usecase.Position, l *logger.Log) {
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
