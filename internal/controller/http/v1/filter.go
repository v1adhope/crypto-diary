// TODO
package v1

import (
	"errors"
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

func getValidMapFilters(c *gin.Context) entity.Filters {
	queryValByQueryKey := entity.Filters{}

	queryVal, err := getAndvalidateDate(c.Query(entity.FilterDate))
	if err == nil {
		queryValByQueryKey[entity.FilterDate] = queryVal
	}

	queryVal, err = getAndvalidatePair(c.Query(entity.FilterPair))
	if err == nil {
		queryValByQueryKey[entity.FilterPair] = queryVal
	}

	queryVal, err = getAndvalidateStrategically(c.Query(entity.FilterStrategically))
	if err == nil {
		queryValByQueryKey[entity.FilterStrategically] = queryVal
	}

	return queryValByQueryKey
}

func getAndvalidateDate(target string) (string, error) {
	_, err := time.Parse(time.DateOnly, target)
	if err != nil {
		return "", errors.New("not valid date")
	}

	return target, nil
}

func getAndvalidatePair(target string) (string, error) {
	if tl := len(target); tl > 12 || tl == 0 {
		return "", errors.New("not valid pair")
	}

	return target, nil
}

func getAndvalidateStrategically(target string) (string, error) {
	_, err := strconv.ParseBool(target)
	if err != nil {
		return "", errors.New("not valid strategically")
	}

	return target, nil
}
