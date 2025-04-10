package timeseries

/*
type Event struct {
	Origin     Origin        `json:"origin"`
	Path       string        `json:"path"` // uri path
	Start      time.Time     `json:"start-ts"`
	Duration   time.Duration `json:"duration"`
	StatusCode int           `json:"status-code"`
}

// This needs to support two metrics for resiliency:
// Threshold - in milliseconds, this is the target for latency, beyond which the service will fail

type Observation struct {
	Origin     Origin  `json:"origin"`
	Uuid       string  `json:"uuid"`
	Saturation float64 `json:"saturation"` // Threshold/Percentile
	Gradiant   float64 `json:"gradiant"`
}

type Action struct {
	Origin Origin `json:"origin"`
	Uuid   string `json:"uuid"`

	RateLimit float64 `json:"rate-limit"`
}

// What is needed to determine the appropriate rate limit in rps
// 1. Peak Threshold - duration in milliseconds where the service will begin to fail.
//             This needs to be a stored value based on the 99th percentile
//             Can this be a database query, returning the high and average for a window??
//             Or can this be a stored value?
//             The problem is, if resiliency is working correctly, this metric is never determined?
//             Also, using the current average would cause excessive rate limiting during off-peak hours
//
//
// 2. Latency   - current 95th percentile for latency for a window
// 3. Gradiant  - current logistic regression gradient for a window


*/
