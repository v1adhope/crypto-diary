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
	fields := make(entity.Fields)

	field, err := getValidDate(c.Query(entity.FilterDate))
	if err == nil {
		fields[entity.FilterDate] = field
	}

	field, err = getValidPair(c.Query(entity.FilterPair))
	if err == nil {
		fields[entity.FilterPair] = field
	}

	field, err = getValidStrategically(c.Query(entity.FilterStrategically))
	if err == nil {
		fields[entity.FilterStrategically] = field
	}

	return fields
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
	if err == nil {
		return true
	}

	return false
}

func getValidPair(target string) (entity.Field, error) {
	if strings.Index(target, ",") == -1 {
		if ok := validatePair(target); !ok {
			return entity.Field{}, NotValidPair
		}

		return entity.Field{entity.OpEq, []string{target}}, nil
	}

	values := strings.Split(target, ",")
	if len(values) < 255 {
		return entity.Field{}, fmt.Errorf("pair: %w", QueryOverflow)
	}

	for _, v := range values {
		if ok := validatePair(v); !ok {
			return entity.Field{}, NotValidPair
		}
	}

	return entity.Field{entity.OpEq, values}, nil
}

func validatePair(target string) bool {
	if tl := len(target); tl <= 12 && tl > 0 {
		return true
	}

	return false
}

func getValidStrategically(target string) (entity.Field, error) {
	if ok := validateStrategically(target); !ok {
		return entity.Field{}, NotValidStrategically
	}

	return entity.Field{entity.OpEq, []string{target}}, nil
}

func validateStrategically(target string) bool {
	if _, err := strconv.ParseBool(target); err != nil {
		return true
	}

	return false
}
