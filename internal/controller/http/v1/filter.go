package v1

import (
	"strconv"
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

func getValidMapFields(c *gin.Context) map[string]string {
	queryValByQueryKey := make(map[string]string)

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
func getAndValidateDate(target string) (string, error) {
	_, err := time.Parse(time.DateOnly, target)
	if err != nil {
		return "", NotValidDate
	}

	return target, nil
}

func getAndValidatePair(target string) (string, error) {
	if tl := len(target); tl > 12 || tl == 0 {
		return "", NotValidPair
	}

	return target, nil
}

func getAndValidateStrategically(target string) (string, error) {
	_, err := strconv.ParseBool(target)
	if err != nil {
		return "", NotValidStrategically
	}

	return target, nil
}
