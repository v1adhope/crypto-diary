package v1

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/controller/http/dto"
)

func registerValidations(v *validator.Validate) {
	v.RegisterValidation("direction", validatePositionDirection)
	v.RegisterValidation("strategically", validatePositionStrategically)
	v.RegisterValidation("risk", validatePositionRisk)
	v.RegisterValidation("deposit", validatePositionDeposit)
	v.RegisterValidation("closePrice", validatePositionClosePrice)
	v.RegisterValidation("positionID", validatePositionID)
	v.RegisterStructValidation(validatePosition, dto.Position{})
}

func validatePositionDirection(fl validator.FieldLevel) bool {
	direction := fl.Field().String()
	if direction == "long" || direction == "short" {
		return true
	}

	return false
}

func validatePositionStrategically(fl validator.FieldLevel) bool {
	_, err := strconv.ParseBool(fl.Field().String())
	if err == nil {
		return true
	}

	return false
}

func validatePositionRisk(fl validator.FieldLevel) bool {
	risk, err := strToFloat(fl.Field().String())
	if err == nil && risk > 0 && risk <= 100 {
		return true
	}

	return false
}

func validatePositionDeposit(fl validator.FieldLevel) bool {
	deposit, err := strToFloat(fl.Field().String())
	if err == nil && deposit > 0 {
		return true
	}

	return false
}

func validatePositionClosePrice(fl validator.FieldLevel) bool {
	tmp := fl.Field().String()
	if tmp == "" {
		return true
	}

	closePrice, err := strToFloat(tmp)
	if err == nil && closePrice >= 0 {
		return true
	}

	return false
}

func validatePositionID(fl validator.FieldLevel) bool {
	_, err := strToFloat(fl.Field().String())
	if err == nil {
		return true
	}

	return false
}

// TODO: Dirty
func validatePosition(sl validator.StructLevel) {
	p := sl.Current().Interface().(dto.Position)

	openPrice, err := strToFloat(p.OpenPrice)
	if err != nil || openPrice < 0 {
		sl.ReportError(p.OpenPrice, "OpenPrice", "", "", "")
	}

	reportStopLoss := func() {
		sl.ReportError(p.StopLossPrice, "StopLossPrice", "", "", "")
	}
	reportTakeProfit := func() {
		sl.ReportError(p.TakeProfitPrice, "TakeProfitPrice", "", "", "")
	}

	stopLossPrice, err := strToFloat(p.StopLossPrice)
	if err != nil || stopLossPrice < 0 {
		reportStopLoss()
	}

	takeProfitPrice, err := strToFloat(p.TakeProfitPrice)
	if err != nil || takeProfitPrice < 0 {
		reportTakeProfit()
	}

	switch p.Direction {
	case "long":
		if stopLossPrice > openPrice {
			reportStopLoss()
		}

		if takeProfitPrice < openPrice {
			reportTakeProfit()
		}
	case "short":
		if stopLossPrice < openPrice {
			reportStopLoss()
		}

		if takeProfitPrice > openPrice {
			reportTakeProfit()
		}
	}
}

func strToFloat(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1, err
	}

	return v, nil
}
