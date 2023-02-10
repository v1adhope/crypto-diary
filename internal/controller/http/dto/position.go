package dto

import (
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type Position struct {
	ID            string  `json:"id,omitempty"`
	OpenDate      string  `json:"openDate" validate:"required,datetime=2006/01/02"`
	Pair          string  `json:"pair" validate:"required,max=12"`
	Reason        string  `json:"reason" validate:"required,max=300"`
	Strategically bool    `json:"strategically"`
	Risk          float64 `json:"risk" validate:"required,gt=0,lt=100"`

	//TODO: direction
	Direction string `json:"direction" validate:"required"`

	//TODO: custom validate
	Deposit         string `json:"deposit" validate:"required"`
	OpenPrice       string `json:"openPrice" validate:"required"`
	StopLossPrice   string `json:"stopLossPrice" validate:"required"`
	TakeProfitPrice string `json:"takeProfitPrice" validate:"required"`
	ClosePrice      string `json:"closePrice"`

	UserID string `json:"-" validate:"required"`
}

func (p *Position) ToEntity() *entity.Position {
	strategically := fmt.Sprintf("%t", p.Strategically)
	risk := fmt.Sprintf("%g", p.Risk)

	return &entity.Position{
		ID:              p.ID,
		OpenDate:        p.OpenDate,
		Pair:            p.Pair,
		Reason:          p.Reason,
		Strategically:   strategically,
		Risk:            risk,
		Direction:       p.Direction,
		Deposit:         p.Deposit,
		OpenPrice:       p.OpenPrice,
		StopLossPrice:   p.StopLossPrice,
		TakeProfitPrice: p.TakeProfitPrice,
		ClosePrice:      p.ClosePrice,
		UserID:          p.UserID,
	}
}

type PositionDelete struct {
	ID string `json:"id"`
}
