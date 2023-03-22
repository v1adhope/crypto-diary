package dto

import (
	"strconv"

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
	return &entity.Position{
		ID:              p.ID,
		OpenDate:        p.OpenDate,
		Pair:            p.Pair,
		Reason:          p.Reason,
		Strategically:   strconv.FormatBool(p.Strategically),
		Risk:            strToFloat(p.Risk),
		Direction:       p.Direction,
		Deposit:         strToFloat(p.Deposit),
		OpenPrice:       strToFloat(p.OpenPrice),
		StopLossPrice:   strToFloat(p.StopLossPrice),
		TakeProfitPrice: strToFloat(p.TakeProfitPrice),
		ClosePrice:      strToFloat(p.ClosePrice),
		UserID:          p.UserID,
	}
}

func strToFloat(s float64) string {
	return strconv.FormatFloat(s, 'f', -1, 64)
}

type PositionDelete struct {
	ID string `json:"id"`
}
