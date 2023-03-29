package v1

import "github.com/gin-gonic/gin"

func catchErrorBind(c *gin.Context, err error) {
	c.Error(err).SetType(gin.ErrorTypeBind)
}

func catchErrorPublic(c *gin.Context, err error) {
	c.Error(err).SetType(gin.ErrorTypePublic)
}
