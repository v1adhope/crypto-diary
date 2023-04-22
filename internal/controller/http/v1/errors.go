package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	NotValidDate          = errors.New("not valid date")
	NotValidPair          = errors.New("not valid pair")
	NotValidStrategically = errors.New("not valid strategically")

	NotValidPositionID = errors.New("invalid id")
)

func catchErrorBind(c *gin.Context, err error) {
	c.Error(err).SetType(gin.ErrorTypeBind)
}

func catchErrorPublic(c *gin.Context, err error) {
	c.Error(err).SetType(gin.ErrorTypePublic)
}
