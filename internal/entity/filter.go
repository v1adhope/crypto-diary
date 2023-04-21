package entity

type Filter struct {
	PaginationCursor int
	Fields           map[string]string
}

const (
	FilterDate          = "date"
	FilterPair          = "pair"
	FilterStrategically = "strategically"
)
