package entity

type (
	Field struct {
		Operation string
		Values    []string
	}

	Filter struct {
		PaginationCursor int
		Fields           map[string]Field
	}
)

const (
	FilterDate          = "date"
	FilterPair          = "pair"
	FilterStrategically = "strategically"

	OpEq    = "eq"
	OpRange = "range"
)
