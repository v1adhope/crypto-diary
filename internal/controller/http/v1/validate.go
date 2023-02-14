package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/controller/http/dto"
)

func validatePositionDirection(fl validator.FieldLevel) bool {
	if fl.Field().String() == "long" || fl.Field().String() == "short" {
		return true
	}

	return false
}

func validatePosition(sl validator.StructLevel) {
	position := sl.Current().Interface().(dto.Position)

	switch position.Direction {
	case "long":
		if position.StopLossPrice > position.OpenPrice {
			sl.ReportError(position.StopLossPrice, "StopLossPrice", "", "", "")
		}

		if position.TakeProfitPrice < position.OpenPrice {
			sl.ReportError(position.TakeProfitPrice, "TakeProfitPrice", "", "", "")
		}
	case "short":
		if position.StopLossPrice < position.OpenPrice {
			sl.ReportError(position.StopLossPrice, "StopLossPrice", "", "", "")
		}

		if position.TakeProfitPrice > position.OpenPrice {
			sl.ReportError(position.TakeProfitPrice, "TakeProfitPrice", "", "", "")
		}
	}
}
