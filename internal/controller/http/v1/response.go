package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, gin.H{"msg": msg})
}

func okResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
