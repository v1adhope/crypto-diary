// TODO: Strong dependency on entity
package v1

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/internal/entity"
)

const (
	_paginationCursorQueryKey = "cursor"
)

func getPaginationCursor(c *gin.Context) int {
	pc, err := strconv.Atoi(c.Query(_paginationCursorQueryKey))
	if err != nil {
		return 0
	}

	return pc
}

func getValidMapFields(c *gin.Context) entity.Fields {
	fieldByFilterName := make(entity.Fields)

	field, err := getValidDate(c.Query(entity.FilterDate))
	if err == nil {
		fieldByFilterName[entity.FilterDate] = field
	}

	field, err = getValidPair(c.Query(entity.FilterPair))
	if err == nil {
		fieldByFilterName[entity.FilterPair] = field
	}

	field, err = getValidStrategically(c.Query(entity.FilterStrategically))
	if err == nil {
		fieldByFilterName[entity.FilterStrategically] = field
	}

	return fieldByFilterName
}

// INFO: Shadow errors, used for debug
func getValidDate(target string) (entity.Field, error) {
	if strings.Index(target, ":") == -1 {
		if ok := validateDate(target); !ok {
			return entity.Field{}, NotValidDate
		}

		return entity.Field{entity.OpEq, []string{target}}, nil
	}

	values := strings.Split(target, ":")
	if len(values) != 2 {
		return entity.Field{}, NotValidDate
	}

	for _, v := range values {
		if ok := validateDate(v); !ok {
			return entity.Field{}, NotValidDate
		}
	}

	return entity.Field{entity.OpRange, values}, nil
}

func validateDate(target string) bool {
	_, err := time.Parse(time.DateOnly, target)
	if err != nil {
		return false
	}

	return true
}

func getValidPair(target string) (entity.Field, error) {
	if len(target) == 255 {
		return entity.Field{}, fmt.Errorf("pair: %w", QueryOverflow)
	}

	if strings.Index(target, ",") == -1 {
		if ok := validatePair(target); !ok {
			return entity.Field{}, NotValidPair
		}

		return entity.Field{entity.OpEq, []string{target}}, nil
	}

	values := strings.Split(target, ",")
	if len(values) == 0 {
		return entity.Field{}, NotValidPair
	}

	for _, v := range values {
		if ok := validatePair(v); !ok {
			return entity.Field{}, NotValidPair
		}
	}

	return entity.Field{entity.OpEq, values}, nil
}

func validatePair(target string) bool {
	if tl := len(target); tl > 12 || tl <= 0 {
		return false
	}

	return true
}

func getValidStrategically(target string) (entity.Field, error) {
	if ok := validateStrategically(target); !ok {
		return entity.Field{}, NotValidStrategically
	}

	return entity.Field{entity.OpEq, []string{target}}, nil
}

func validateStrategically(target string) bool {
	if _, err := strconv.ParseBool(target); err != nil {
		return false
	}

	return true
}
