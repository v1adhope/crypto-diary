package entity

type Position struct {
	ID              string `json:"id"`
	OpenDate        string `json:"openDate"`
	Pair            string `json:"pair"`
	Reason          string `json:"reason"`
	Strategically   string `json:"strategically"`
	Risk            string `json:"risk"`
	Direction       string `json:"direction"`
	Deposit         string `json:"deposit"`
	OpenPrice       string `json:"openPrice"`
	StopLossPrice   string `json:"stopLossPrice"`
	TakeProfitPrice string `json:"takeProfitPrice"`
	ClosePrice      string `json:"closePrice"`
	UserID          string `json:"userID"`
}
