package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	_userCtx = "userID"
)

func (r *Router) AuthorizeJWT() gin.HandlerFunc {
	const _bearerSchema = "Bearer "

	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
			return
		}

		clientToken := header[len(_bearerSchema):]

		id, err := r.UseCases.User.CheckAuth(c.Request.Context(), clientToken)
		if err != nil {
			r.Logger.Debug(err, "http/v1: AutorizeJWT")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
			return
		}

		c.Set(_userCtx, id)

		c.Next()
	}
}
