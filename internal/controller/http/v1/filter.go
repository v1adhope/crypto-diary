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

type Field struct {
	Operation string
	Values    []string
}

func getValidMapFields(c *gin.Context) map[string]entity.Field {
	queryValByQueryKey := make(map[string]entity.Field)

	queryVal, err := getAndValidateDate(c.Query(entity.FilterDate))
	if err == nil {
		queryValByQueryKey[entity.FilterDate] = queryVal
	}

	queryVal, err = getAndValidatePair(c.Query(entity.FilterPair))
	if err == nil {
		queryValByQueryKey[entity.FilterPair] = queryVal
	}

	queryVal, err = getAndValidateStrategically(c.Query(entity.FilterStrategically))
	if err == nil {
		queryValByQueryKey[entity.FilterStrategically] = queryVal
	}

	return queryValByQueryKey
}

// INFO: Shadow errors, used for debug
func getAndValidateDate(target string) (entity.Field, error) {
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

func getAndValidatePair(target string) (entity.Field, error) {
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

func getAndValidateStrategically(target string) (entity.Field, error) {
	_, err := strconv.ParseBool(target)
	if err != nil {
		return entity.Field{}, NotValidStrategically
	}

	return entity.Field{entity.OpEq, []string{target}}, nil
}
