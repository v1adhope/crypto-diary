package entity

type (
	Field struct {
		Operation string
		Values    []string
	}

	Fields map[string]Field

	Filter struct {
		PaginationCursor int
		Fields           Fields
	}
)

const (
	FilterDate          = "date"
	FilterPair          = "pair"
	FilterStrategically = "strategically"

	OpEq    = "eq"
	OpRange = "range"
)
