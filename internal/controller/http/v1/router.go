package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/hash"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

type Deps struct {
	Handler  *gin.Engine
	UseCases *usecase.UseCases
	Logger   *logger.Logger
	Validate *validator.Validate
	Hasher   hash.PasswordHasher
}

// TODO
func NewRouter(d *Deps) {
	d.Handler.Use(gin.Logger(), gin.Recovery())

	h := d.Handler.Group("/v1")
	{
		newUserRoutes(&userRoutes{
			h:        h,
			l:        d.Logger,
			validate: d.Validate,
			useCase:  d.UseCases.User,
			hasher:   d.Hasher,
		})

		// newPositionRoutes(h, d.UseCases.Position, d.Logger)
	}
}
