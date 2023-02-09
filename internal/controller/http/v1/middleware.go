package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Extraxt id
func (r *Router) AuthorizeJWT() gin.HandlerFunc {
	const _bearerSchema = "Bearer "

	return func(c *gin.Context) {
		header := c.GetHeader("Autorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
			return
		}

		clientToken := header[len(_bearerSchema):]

		err := r.UseCases.User.CheckAuth(c.Request.Context(), clientToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
			return
		}

		c.Next()
	}
}
