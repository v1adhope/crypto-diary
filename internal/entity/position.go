package entity

import (
	"fmt"
	"strconv"
)

type Position struct {
	ID              string `json:"ID"`
	OpenDate        string `json:"openDate"`
	Pair            string `json:"pair"`
	Reason          string `json:"reason"`
	AccordingToPlan string `json:"accordingToPlan"`
	Risk            string `json:"risk"`
	Direction       string `json:"direction"`
	Deposit         string `json:"deposit"`
	OpenPrice       string `json:"openPrice"`
	StopLossPrice   string `json:"stopLossPrice"`
	TakeProfitPrice string `json:"takeProfitPrice"`
	ClosePrice      string `json:"closePrice"`
	UserID          string `json:"userID"`
}

func (u *Position) ValidPosition() error {
	var errBuff error

	if len(u.Pair) > 12 || len(u.Pair) < 3 {
		errBuff = fmt.Errorf("ticker does not exist:")
	}

	if u.Reason == "" {
		errBuff = fmt.Errorf("%s specify the reason:", errBuff)
	}

	_, err := strconv.ParseBool(u.AccordingToPlan)
	if err != nil {
		errBuff = fmt.Errorf("%s according to plan true of false:", errBuff)
	}

	risk, err := strconv.ParseFloat(u.Risk, 64)
	if err != nil || risk < 0 || risk > 100 {
		errBuff = fmt.Errorf("%s impossible risk:", errBuff)
	}

	switch u.Direction {
	default:
		errBuff = fmt.Errorf("%s unknown direction:", errBuff)
	case "long":
		if u.StopLossPrice > u.OpenPrice {
			errBuff = fmt.Errorf("%s stop loss cannot be greater than the open price:", errBuff)
		}
		if u.TakeProfitPrice < u.OpenPrice {
			errBuff = fmt.Errorf("%s take profit cannot be less than the open price:", errBuff)
		}
	case "short":
		if u.StopLossPrice < u.OpenPrice {
			errBuff = fmt.Errorf("%s stop loss cannot be less than the open price:", errBuff)
		}
		if u.TakeProfitPrice > u.OpenPrice {
			errBuff = fmt.Errorf("%s take profit cannot be greater than the open price:", errBuff)
		}
	}

	deposit, err := strconv.ParseUint(u.Deposit, 10, 64)
	if err != nil || deposit == 0 {
		errBuff = fmt.Errorf("%s position deposit makes no sense:", errBuff)
	}

	openPrice, err := strconv.Atoi(u.OpenPrice)
	if err != nil || openPrice <= 0 {
		errBuff = fmt.Errorf("%s imposible open price:", errBuff)
	}

	closePrice, err := strconv.Atoi(u.OpenPrice)
	if err != nil || closePrice <= 0 {
		errBuff = fmt.Errorf("%s imposible close price:", errBuff)
	}

	if err != nil {
		return errBuff
	}
	return nil
}

// TODO: remove or replace
// type CreatePositionDTO struct {
// 	OpenDate        string `json:"openDate"`
// 	Pair            string `json:"pair"`
// 	PercentageRisk  string `json:"percentageRisk"`
// 	Direction       string `json:"direction"`
// 	Deposit         string `json:"deposit"`
// 	OpenPrice       string `json:"openPrice"`
// 	StopLossPrice   string `json:"stopLossPrice"`
// 	TakeProfitPrice string `json:"takeProfitPrice"`
// 	ClosePrice      string `json:"closePrice"`
// 	UserID          string `json:"userID"`
// }
