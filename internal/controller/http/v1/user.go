package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/hash"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

// TODO
type userRoutes struct {
	h        *gin.RouterGroup
	l        *logger.Logger
	validate *validator.Validate
	useCase  usecase.User
	hasher   hash.PasswordHasher
}

func newUserRoutes(r *userRoutes) {
	h := r.h.Group("/auth")
	{
		h.POST("/signup", r.signUp)
		h.POST("/signin", r.signIn)
	}
}

type signRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *userRoutes) signUp(c *gin.Context) {
	request := &signRequest{}

	if err := c.BindJSON(&request); err != nil {
		r.l.Logger.Err(err).Msg("http/v1: signUp: BindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})

		return
	}

	if err := r.validate.Struct(request); err != nil {
		r.l.Logger.Err(err).Msg("http/v1: signUp: validate")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})

		return
	}

	encryptedPassword, err := r.hasher.GenerateEncryptedPassword(request.Password)
	if err != nil {
		r.l.Logger.Err(err).Msg("http/v1: signUp: GenerateEncryptedPassword")
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	if err := r.useCase.SignUp(c.Request.Context(), request.Email, encryptedPassword); err != nil {
		r.l.Logger.Err(err).Msg("http/v1: signUp: useCase")

		err = errors.Unwrap(err)

		if err == entity.ErrUserAlreadyExists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s", err)})

			return

		}

		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.Status(http.StatusOK)
}

type signResponse struct {
	ID string `json:"id"`
}

func (ur *userRoutes) signIn(c *gin.Context) {

}
