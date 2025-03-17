package timeseries

import "time"

type Event struct {
	Start    time.Time
	Duration int // milliseconds
	
}
