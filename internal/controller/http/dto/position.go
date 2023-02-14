package dto

import (
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type Position struct {
	ID              string  `json:"id,omitempty"`
	OpenDate        string  `json:"openDate" validate:"required,datetime=2006/01/02"`
	Pair            string  `json:"pair" validate:"required,max=12"`
	Reason          string  `json:"reason" validate:"required,max=300"`
	Strategically   bool    `json:"strategically"`
	Risk            float64 `json:"risk" validate:"required,gt=0,lt=100"`
	Direction       string  `json:"direction" validate:"required,direction"`
	Deposit         float64 `json:"deposit" validate:"required,gt=0"`
	OpenPrice       float64 `json:"openPrice" validate:"required,gt=0"`
	StopLossPrice   float64 `json:"stopLossPrice" validate:"required,gt=0"`
	TakeProfitPrice float64 `json:"takeProfitPrice" validate:"required,gt=0"`
	ClosePrice      float64 `json:"closePrice" validate:"gt=0"`
	UserID          string  `json:"-" validate:"required"`
}

func (p *Position) ToEntity() *entity.Position {
	strategically := fmt.Sprintf("%t", p.Strategically)
	risk := fmt.Sprintf("%g", p.Risk)
	deposit := fmt.Sprintf("%g", p.Deposit)
	openPrice := fmt.Sprintf("%g", p.OpenPrice)
	stopLossPrice := fmt.Sprintf("%g", p.StopLossPrice)
	takeProfitPrice := fmt.Sprintf("%g", p.TakeProfitPrice)
	closePrice := fmt.Sprintf("%g", p.ClosePrice)

	return &entity.Position{
		ID:              p.ID,
		OpenDate:        p.OpenDate,
		Pair:            p.Pair,
		Reason:          p.Reason,
		Strategically:   strategically,
		Risk:            risk,
		Direction:       p.Direction,
		Deposit:         deposit,
		OpenPrice:       openPrice,
		StopLossPrice:   stopLossPrice,
		TakeProfitPrice: takeProfitPrice,
		ClosePrice:      closePrice,
		UserID:          p.UserID,
	}
}

type PositionDelete struct {
	ID string `json:"id"`
}
