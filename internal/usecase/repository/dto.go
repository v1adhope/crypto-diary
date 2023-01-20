package repository

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/v1adhope/crypto-diary/internal/entity"
)

const timeModel = "2006-01-02"

type PositionDTO struct {
	ID              string
	OpenDate        pgtype.Date
	Pair            string
	Reason          string
	AccordingToPlan bool
	Risk            string
	Direction       string
	Deposit         string
	OpenPrice       string
	StopLossPrice   string
	TakeProfitPrice string
	ClosePrice      pgtype.Text
	UserID          string
}

func (pd *PositionDTO) ToEntity() *entity.Position {
	position := &entity.Position{}

	position.ID = pd.ID
	position.Pair = pd.Pair
	position.Reason = pd.Reason
	position.Risk = pd.Risk
	position.Direction = pd.Direction
	position.Deposit = pd.Deposit
	position.OpenPrice = pd.OpenPrice
	position.StopLossPrice = pd.StopLossPrice
	position.TakeProfitPrice = pd.TakeProfitPrice
	position.UserID = pd.UserID

	position.OpenDate = fmt.Sprintf("%s", pd.OpenDate.Time.Format(timeModel))
	position.AccordingToPlan = fmt.Sprintf("%t", pd.AccordingToPlan)
	if pd.ClosePrice.Valid {
		position.ClosePrice = fmt.Sprintf("%s", pd.ClosePrice.String)
	}

	return position
}
