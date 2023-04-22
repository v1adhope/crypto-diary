package repository

import (
	"sync"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

var allowedFilters sync.Map

func init() {
	allowedFilters.Store(entity.FilterDate, "open_date")
	allowedFilters.Store(entity.FilterPair, "pair")
	allowedFilters.Store(entity.FilterStrategically, "strategically")
}
