package repository

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/v1adhope/crypto-diary/internal/entity"
)

const _timeModel = "2006-01-02"

type PositionDTO struct {
	ID              string
	OpenDate        pgtype.Date
	Pair            string
	Reason          pgtype.Text
	Strategically   bool
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
	opendate := pd.OpenDate.Time.Format(_timeModel)
	strategically := strconv.FormatBool(pd.Strategically)

	reason := nullTextCheck(&pd.Reason)
	closePrice := nullTextCheck(&pd.ClosePrice)

	return &entity.Position{
		ID:              pd.ID,
		OpenDate:        opendate,
		Pair:            pd.Pair,
		Reason:          reason,
		Strategically:   strategically,
		Risk:            pd.Risk,
		Direction:       pd.Direction,
		Deposit:         pd.Deposit,
		OpenPrice:       pd.OpenPrice,
		StopLossPrice:   pd.StopLossPrice,
		TakeProfitPrice: pd.TakeProfitPrice,
		ClosePrice:      closePrice,
		UserID:          pd.UserID,
	}
}

func nullTextCheck(t *pgtype.Text) string {
	if t.Valid {
		return t.String
	}

	return ""
}
