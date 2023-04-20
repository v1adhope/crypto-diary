// TODO: Params struct, names, clean
package repository

import (
	"fmt"
	"strings"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type filterDeps struct {
	Query                 string
	QueryPlaceholderCount int
	PaginationCursor      int
	Filters               entity.Filters
	AllowedFilters        entity.Filters
}

func BuildFilterString(deps filterDeps) (string, []string) {
	var (
		filterRaw strings.Builder
	)

	args, argsCounter := make([]string, 0), deps.QueryPlaceholderCount+1

	for k, v := range deps.Filters {
		if realFilterName, ok := deps.AllowedFilters[k]; ok {
			fmt.Fprintf(&filterRaw, "AND %s = $%d ", realFilterName, argsCounter)
			args = append(args, v)
			argsCounter++
		}
	}

	fmt.Fprintf(&filterRaw, "ORDER by position_id ASC LIMIT $%d", argsCounter)

	return fmt.Sprintf(deps.Query, filterRaw.String()), args
}
