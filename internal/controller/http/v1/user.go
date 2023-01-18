package v1

import "github.com/gin-gonic/gin"

// TODO
type userRoutes struct {
}

func newUserRoutes(handler *gin.RouterGroup) {

	h := handler.Group("/auth")
	{
		h.POST("/sign-up")
		h.POST("/sign-in")
	}
}
