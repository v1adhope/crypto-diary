package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	_userCtxKey = "userID"
)

type Middleware interface {
	AuthorizeJWT() gin.HandlerFunc
}

func (r *Router) AuthorizeJWT() gin.HandlerFunc {
	const _bearerSchema = "Bearer"

	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
			return
		}

		clientToken := strings.Split(header, " ")
		if len(clientToken) != 2 || clientToken[0] != _bearerSchema {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
			return
		}

		id, err := r.UseCases.User.CheckAuth(c.Request.Context(), clientToken[1])
		if err != nil {
			r.Logger.Debug(err, "http/v1: AutorizeJWT")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
			return
		}

		c.Set(_userCtxKey, id)

		c.Next()
	}
}
