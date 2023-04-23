package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/controller/http/dto"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

const (
	_positionIDQueryKey = "id"
)

type positionRoutes struct {
	h        *gin.RouterGroup
	m        authMiddleware
	validate *validator.Validate
	useCase  usecase.Position
	logger   logger.Logger
}

func newPositionRoutes(r *positionRoutes) {
	h := r.h.Group("/position")

	h.Use(r.m.tokenHandler())
	{
		h.GET("/", r.GetAll)
		h.POST("/", r.Create)
		h.DELETE("/", r.Delete)
		h.PUT("/", r.Replace)
	}

	registerValidations(r.validate)
}

func (r *positionRoutes) GetAll(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	filter := entity.Filter{
		PaginationCursor: getPaginationCursor(c),
		Fields:           getValidMapFields(c),
	}

	positions, err := r.useCase.GetAll(c.Request.Context(), userID, filter)
	if err != nil {
		r.logger.Debug(err, "http/v1: GetAll position: GetAll")
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, positions)
}

func (r *positionRoutes) Create(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	positionDTO := &dto.Position{
		UserID: userID,
	}

	if err := c.ShouldBindJSON(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Create position: ShouldBindJSON")
		catchErrorBind(c, err)
		return
	}

	if err := r.validate.Struct(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Create position: Struct")
		catchErrorPublic(c, err)
		return
	}

	position := positionDTO.ToEntity()

	err = r.useCase.Create(c.Request.Context(), position)
	if err != nil {
		r.logger.Debug(err, "http/v1: Create position: Create")
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": position.ID,
	})
}

func (r *positionRoutes) Delete(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	dto := &dto.PositionDelete{
		ID: c.Query(_positionIDQueryKey),
	}

	if err := r.validate.Struct(dto); err != nil {
		r.logger.Debug(err, "http/v1: Delete position: Struct")
		catchErrorPublic(c, NotValidPositionID)
		return
	}

	if err := r.useCase.Delete(c.Request.Context(), userID, dto.ID); err != nil {
		r.logger.Debug(err, "http/v1: Delete position: Delete")
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *positionRoutes) Replace(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	positionDTO := &dto.PositionUpdate{
		Position: &dto.Position{UserID: userID},
	}

	if err := c.ShouldBindJSON(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Replace position: ShouldBindJSON")
		catchErrorBind(c, err)
		return
	}

	if err := r.validate.Struct(positionDTO); err != nil {
		r.logger.Debug(err, "http/v1: Replace position: Struct")
		catchErrorPublic(c, err)
		return
	}

	position := positionDTO.ToEntity()

	if err := r.useCase.Replace(c.Request.Context(), position); err != nil {
		r.logger.Debug(err, "http/v1: Replace position: Replace")
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
