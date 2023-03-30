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
	v.RegisterStructValidation(validatePosition, dto.Position{})
}

func validatePositionDirection(fl validator.FieldLevel) bool {
	direction := fl.Field().String()
	if direction == "long" || direction == "short" {
		return true
	}

	return false
}

func validatePositionStrategically(fls validator.FieldLevel) bool {
	_, err := strconv.ParseBool(fls.Field().String())
	if err == nil {
		return true
	}

	return false
}

func validatePositionRisk(fls validator.FieldLevel) bool {
	risk, err := strToFloat(fls.Field().String())
	if err == nil && risk > 0 && risk <= 100 {
		return true
	}

	return false
}

func validatePositionDeposit(fls validator.FieldLevel) bool {
	deposit, err := strToFloat(fls.Field().String())
	if err == nil && deposit > 0 {
		return true
	}

	return false
}

func validatePositionClosePrice(fls validator.FieldLevel) bool {
	tmp := fls.Field().String()
	if tmp == "" {
		return true
	}

	closePrice, err := strToFloat(tmp)
	if err == nil && closePrice >= 0 {
		return true
	}

	return false
}

func validatePosition(sl validator.StructLevel) {
	p := sl.Current().Interface().(dto.Position)

	openPrice, err := strToFloat(p.OpenPrice)
	if err != nil || openPrice < 0 {
		sl.ReportError(p.OpenPrice, "OpenPrice", "", "", "")
	}

	stopLossPrice, err := strToFloat(p.StopLossPrice)
	if err != nil || stopLossPrice < 0 {
		sl.ReportError(p.StopLossPrice, "StopLossPrice", "", "", "")
	}

	takeProfitPrice, err := strToFloat(p.TakeProfitPrice)
	if err != nil || takeProfitPrice < 0 {
		sl.ReportError(p.TakeProfitPrice, "TakeProfitPrice", "", "", "")
	}

	switch p.Direction {
	case "long":
		if stopLossPrice > openPrice {
			sl.ReportError(p.StopLossPrice, "StopLossPrice", "", "", "")
		}

		if takeProfitPrice < openPrice {
			sl.ReportError(p.TakeProfitPrice, "TakeProfitPrice", "", "", "")
		}
	case "short":
		if stopLossPrice < openPrice {
			sl.ReportError(p.StopLossPrice, "StopLossPrice", "", "", "")
		}

		if takeProfitPrice > openPrice {
			sl.ReportError(p.TakeProfitPrice, "TakeProfitPrice", "", "", "")
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
