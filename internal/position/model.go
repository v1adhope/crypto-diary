package position

import "github.com/v1adhope/crypto-diary/internal/user"

type Position struct {
	PositionID      string    `json:"positionID"`
	OpenDate        string    `json:"openDate"`
	Pair            string    `json:"pair"`
	PercentageRisk  string    `json:"percentageRisk"`
	Direction       string    `json:"direction"`
	Deposit         string    `json:"deposit"`
	OpenPrice       string    `json:"openPrice"`
	StopLossPrice   string    `json:"stopLossPrice"`
	TakeProfitPrice string    `json:"takeProfitPrice"`
	ClosePrice      string    `json:"closePrice"`
	User            user.User `json:"user"`
}
