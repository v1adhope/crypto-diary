package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/controller/http/dto"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

type userRoutes struct {
	h        *gin.RouterGroup
	validate *validator.Validate
	useCase  usecase.User
	logger   logger.Logger
}

func newUserRoutes(r *userRoutes) {
	h := r.h.Group("/auth")
	{
		h.POST("/signup", r.signUp)
		h.POST("/signin", r.signIn)
		h.POST("/refresh", r.refreshTokens)
		h.POST("/signout", r.signOut)
	}
}

func (r *userRoutes) signUp(c *gin.Context) {
	request := &dto.SignRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		r.logger.Debug(err, "http/v1: signUp: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	if err := r.validate.Struct(request); err != nil {
		r.logger.Debug(err, "http/v1: signUp: Struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	if err := r.useCase.SignUp(c.Request.Context(), request.Email, request.Password); err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			r.logger.Debug(err, "http/v1: signUp: SignUp")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": entity.ErrUserAlreadyExists.Error()})
			return
		}

		r.logger.Error(err, "http/v1: signUp: SignUp")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (r *userRoutes) signIn(c *gin.Context) {
	request := &dto.SignRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		r.logger.Debug(err, "http/v1: signIn: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	if err := r.validate.Struct(request); err != nil {
		r.logger.Debug(err, "http/v1: signIn: Struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	refreshToken, accessToken, err := r.useCase.SignIn(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) {
			r.logger.Debug(err, "http/v1: signIn: SignIn")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": entity.ErrUserNotExists.Error()})
			return
		}

		if errors.Is(err, entity.ErrWrongPassword) {
			r.logger.Debug(err, "http/v1: signIn: SignIn")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
			return
		}

		r.logger.Error(err, "http/v1: signIn: SignIn")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}

func (r *userRoutes) refreshTokens(c *gin.Context) {
	clientToken := &dto.RefreshToken{}

	if err := c.ShouldBindJSON(clientToken); err != nil {
		r.logger.Debug(err, "http/v1: refreshTokens: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid token"})
		return
	}

	refreshToken, accessToken, err := r.useCase.ReissueTokens(c.Request.Context(), clientToken.Token)
	if err != nil {
		r.logger.Debug(err, "http/v1: refreshTokens: ReissueTokens")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}

func (r *userRoutes) signOut(c *gin.Context) {
	clientToken := &dto.RefreshToken{}

	if err := c.ShouldBindJSON(clientToken); err != nil {
		r.logger.Debug(err, "http/v1: signOut: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid token"})
		return
	}

	if err := r.useCase.SignOut(c.Request.Context(), clientToken.Token); err != nil {
		r.logger.Error(err, "http/v1: signOut: SignOut")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
		return
	}

	c.Status(http.StatusOK)
}
