package v1

import "github.com/gin-gonic/gin"

// TODO
type positionRoutes struct {
}

func newPositionRoutes(handler *gin.RouterGroup) {

	h := handler.Group("/positon")
	{
		h.GET("/")
		h.GET("/:id")
		h.POST("/")
		h.PUT("/")
		h.DELETE("/")
	}
}
