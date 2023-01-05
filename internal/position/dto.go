package position

type CreatePositionDTO struct {
	OpenDate        string `json:"openDate"`
	Pair            string `json:"pair"`
	PercentageRisk  string `json:"percentageRisk"`
	Direction       string `json:"direction"`
	Deposit         string `json:"deposit"`
	OpenPrice       string `json:"openPrice"`
	StopLossPrice   string `json:"stopLossPrice"`
	TakeProfitPrice string `json:"takeProfitPrice"`
	ClosePrice      string `json:"closePrice"`
	UserID          string `json:"userID"`
}
