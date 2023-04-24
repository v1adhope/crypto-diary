package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/internal/entity"
)

const (
	ctxKeyUser = "userID"
)

type authMiddleware interface {
	tokenHandler() gin.HandlerFunc
}

func (r *Router) tokenHandler() gin.HandlerFunc {
	const bearerSchema = "Bearer"

	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.Error(entity.ErrTokenInvalid)
			return
		}

		clientToken := strings.Split(header, " ")
		if len(clientToken) != 2 || clientToken[0] != bearerSchema {
			c.Error(entity.ErrTokenInvalid)
			return
		}

		id, err := r.UseCases.User.CheckAuth(c.Request.Context(), clientToken[1])
		if err != nil {
			r.Logger.Debug(err, "http/v1: tokenHandler: CheckAuth")
			c.Error(entity.ErrTokenInvalid)
			return
		}

		c.Set(ctxKeyUser, id)

		c.Next()
	}
}

func getUserID(c *gin.Context) (string, error) {
	userID := c.GetString(ctxKeyUser)
	if userID == "" {
		return "", entity.ErrTokenInvalid
	}

	return userID, nil
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			if err.IsType(gin.ErrorTypeBind) {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "invalid JSON",
				})
				return
			}

			if err.IsType(gin.ErrorTypePublic) {
				if errors.Is(err, NotValidPositionID) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": NotValidPositionID.Error(),
					})
					return
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "invalid data",
				})
				return
			}

			if err.IsType(gin.ErrorTypePrivate) {
				//INFO: User
				if errors.Is(err, entity.ErrUserAlreadyExists) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": entity.ErrUserAlreadyExists.Error(),
					})
					return
				}

				if errors.Is(err, entity.ErrUserNotExists) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": entity.ErrUserNotExists.Error(),
					})
					return
				}

				if errors.Is(err, entity.ErrWrongPassword) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": "incorrect login or password",
					})
					return
				}

				if errors.Is(err, entity.ErrTokenInvalid) || errors.Is(err, entity.ErrTokenInTheBlocklisk) {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": entity.ErrTokenInvalid.Error(),
					})
					return
				}

				//INFO: Positon
				if errors.Is(err, entity.ErrNoFoundPosition) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": entity.ErrNoFoundPosition.Error(),
					})
					return
				}

				if errors.Is(err, entity.ErrNothingToChange) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": entity.ErrNothingToChange.Error(),
					})
					return
				}
			}

			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}
