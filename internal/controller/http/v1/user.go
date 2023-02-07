package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

type userRoutes struct {
	h        *gin.RouterGroup
	l        *logger.Logger
	validate *validator.Validate
	useCase  usecase.User
}

func newUserRoutes(r *userRoutes) {
	h := r.h.Group("/auth")
	{
		h.POST("/signup", r.signUp)
		h.POST("/signin", r.signIn)
		h.POST("/refresh", r.refreshTokens)
	}
}

type signRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=32,min=8"`
}

func (r *userRoutes) signUp(c *gin.Context) {
	request := &signRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		r.l.Logger.Warn().Err(err).Msg("http/v1: signUp: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	if err := r.validate.Struct(request); err != nil {
		r.l.Logger.Warn().Err(err).Msg("http/v1: signUp: Struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	if err := r.useCase.SignUp(c.Request.Context(), request.Email, request.Password); err != nil {
		r.l.Logger.Err(err).Msg("http/v1: signUp: SignUp")

		if errors.Is(err, entity.ErrUserAlreadyExists) {
			r.l.Logger.Warn().Err(err).Msg("http/v1: signUp: SignUp")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": entity.ErrUserAlreadyExists.Error()})
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (r *userRoutes) signIn(c *gin.Context) {
	request := &signRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		r.l.Logger.Warn().Err(err).Msg("http/v1: signIn: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	if err := r.validate.Struct(request); err != nil {
		r.l.Logger.Warn().Err(err).Msg("http/v1: signIn: Struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
		return
	}

	refreshToken, accessToken, err := r.useCase.SignIn(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) {
			r.l.Logger.Warn().Err(err).Msg("http/v1: signIn: SignIn")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": entity.ErrUserNotExists.Error()})
			return
		}

		if errors.Is(err, entity.ErrWrongPassword) {
			r.l.Logger.Warn().Err(err).Msg("http/v1: signIn: SignIn")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "incorrect login or password"})
			return
		}

		r.l.Logger.Err(err).Msg("http/v1: signIn: SignIn")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}

type refreshToken struct {
	Token string `json:"refreshToken"`
}

func (r *userRoutes) refreshTokens(c *gin.Context) {
	clientToken := &refreshToken{}

	if err := c.ShouldBindJSON(clientToken); err != nil {
		r.l.Logger.Warn().Err(err).Msg("http/v1: refreshTokens: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid token"})
		return
	}

	refreshToken, accessToken, err := r.useCase.ReissueTokens(c.Request.Context(), clientToken.Token)
	if err != nil {
		r.l.Logger.Warn().Err(err).Msg("http/v1: refreshTokens: ReissueTokens")
		errorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}
