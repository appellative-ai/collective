package timeseries

import "time"

type Event struct {
	Origin     Origin        `json:"origin"`
	Start      time.Time     `json:"start-ts"`
	Duration   time.Duration `json:"duration"`
	StatusCode int           `json:"status-code"`
}

/*
Cached     bool `json:"cached"`
StatusCode int  `json:"status-code"`
//RequestId string `json:"request-id"`

*/
