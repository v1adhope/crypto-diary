// TODO: Best place?
package entity

type Filters map[string]string

var (
	FilterDate          = "date"
	FilterPair          = "pair"
	FilterStrategically = "strategically"

	AllowedFilters = map[string]string{
		FilterDate:          "open_date",
		FilterPair:          "pair",
		FilterStrategically: "strategically",
	}
)
