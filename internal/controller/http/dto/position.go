package dto

import (
	"github.com/v1adhope/crypto-diary/internal/entity"
)

type Position struct {
	ID              string `json:"id,omitempty"`
	OpenDate        string `json:"openDate" validate:"required,datetime=2006-01-02"`
	Pair            string `json:"pair" validate:"required,max=12"`
	Reason          string `json:"reason" validate:"required,max=300"`
	Strategically   string `json:"strategically" validate:"required,strategically"`
	Risk            string `json:"risk" validate:"required,risk"`
	Direction       string `json:"direction" validate:"required,direction"`
	Deposit         string `json:"deposit" validate:"required,deposit"`
	OpenPrice       string `json:"openPrice" validate:"required"`
	StopLossPrice   string `json:"stopLossPrice" validate:"required"`
	TakeProfitPrice string `json:"takeProfitPrice" validate:"required"`
	ClosePrice      string `json:"closePrice" validate:"closePrice"`
	UserID          string `json:"-" validate:"required"`
}

func (p *Position) ToEntity() *entity.Position {
	return &entity.Position{
		ID:              p.ID,
		OpenDate:        p.OpenDate,
		Pair:            p.Pair,
		Reason:          p.Reason,
		Strategically:   p.Strategically,
		Risk:            p.Risk,
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
