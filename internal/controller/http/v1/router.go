package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

type Router struct {
	Handler  *gin.Engine
	UseCases *usecase.UseCases
	Logger   *logger.Log
	Validate *validator.Validate
}

func NewRouter(r *Router) {
	r.Handler.Use(
		gin.Logger(),
		gin.Recovery(),
	)
	r.Handler.SetTrustedProxies(nil)

	h := r.Handler.Group("/v1")

	h.Use(errorHandler())
	{
		newUserRoutes(&userRoutes{
			h:        h,
			logger:   r.Logger,
			validate: r.Validate,
			useCase:  r.UseCases.User,
		})

		newPositionRoutes(&positionRoutes{
			h:        h,
			m:        r,
			logger:   r.Logger,
			validate: r.Validate,
			useCase:  r.UseCases.Position,
		})
	}
}
