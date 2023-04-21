package repository

import (
	"fmt"
	"strings"
	"sync"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

var allowedFilters = map[string]string{
	entity.FilterDate:          "open_date",
	entity.FilterPair:          "pair",
	entity.FilterStrategically: "strategically",
}

type filterBuilderDeps struct {
	Query  string
	Filter entity.Filter
}

func BuildFilterString(deps filterBuilderDeps) (string, []string) {
	var (
		filterRaw strings.Builder
		mu        sync.RWMutex
	)

	fmt.Fprint(&filterRaw, deps.Query, " ")

	argsCounter := strings.Count(deps.Query, "$") + 1

	args := make([]string, 0)

	for fieldKey, fieldVal := range deps.Filter.Fields {
		mu.RLock()
		if realFilterName, ok := allowedFilters[fieldKey]; ok {
			fmt.Fprintf(&filterRaw, "AND %s = $%d ", realFilterName, argsCounter)
			args = append(args, fieldVal)
			argsCounter++
		}
		mu.RUnlock()
	}

	fmt.Fprintf(&filterRaw, "ORDER by position_id ASC LIMIT $%d", argsCounter)

	return filterRaw.String(), args
}
