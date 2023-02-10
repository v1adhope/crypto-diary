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

type positionRoutes struct {
	h        *gin.RouterGroup
	validate *validator.Validate
	useCase  usecase.Position
	logger   logger.Logger
}

func newPositionRoutes(r *positionRoutes) {
	h := r.h
	{
		h.GET("/", r.GetAll)
		h.POST("/", r.Create)
		h.DELETE("/", r.Delete)
		h.PUT("/", r.Replace)
	}
}

func (r *positionRoutes) GetAll(c *gin.Context) {
	id := c.GetString(_userCtx)

	positions, err := r.useCase.GetAll(c.Request.Context(), id)
	if err != nil {
		r.logger.Debug(err, "http/v1: GetAll position: GetAll")
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, positions)
}

func (r *positionRoutes) Create(c *gin.Context) {
	userID := c.GetString(_userCtx)

	positionDTO := &dto.Position{
		UserID: userID,
	}

	if err := c.ShouldBindJSON(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Create position: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid position"})
		return
	}

	if err := r.validate.Struct(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Create position: Struct")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid position"})
		return
	}

	position := positionDTO.ToEntity()

	err := r.useCase.Create(c.Request.Context(), position)
	if err != nil {
		r.logger.Debug(err, "http/v1: Create position: Create")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, position.ID)
}

func (r *positionRoutes) Delete(c *gin.Context) {
	userID := c.GetString(_userCtx)

	positionDTO := &dto.PositionDelete{}

	if err := c.ShouldBindJSON(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Delete position: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid position id"})
		return
	}

	if err := r.useCase.Delete(c.Request.Context(), userID, positionDTO.ID); err != nil {
		if errors.Is(err, entity.ErrNoFoundPosition) {
			r.logger.Debug(err, "http/v1: Delete position: Delete")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": entity.ErrNoFoundPosition.Error()})
			return
		}

		r.logger.Debug(err, "http/v1: Delete position: Delete")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (r *positionRoutes) Replace(c *gin.Context) {
	userID := c.GetString(_userCtx)

	positionDTO := &dto.Position{
		UserID: userID,
	}

	if err := c.ShouldBindJSON(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Replace position: ShouldBindJSON")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid position"})
		return
	}

	if err := r.validate.Struct(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Replace position: Struct")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid position"})
		return
	}

	position := positionDTO.ToEntity()

	if err := r.useCase.Replace(c.Request.Context(), position); err != nil {
		if errors.Is(err, entity.ErrNothingToChange) {
			r.logger.Debug(err, "http/v1: Replace position: Replace")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": entity.ErrNothingToChange.Error()})
			return
		}

		r.logger.Debug(err, "http/v1: Replace position: Replace")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
