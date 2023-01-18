package v1

import "github.com/gin-gonic/gin"

// TODO
func NewRouter(handler *gin.Engine) {
	h := handler.Group("/v1")
	{
		newPositionRoutes(h)
		newUserRoutes(h)
	}
}
