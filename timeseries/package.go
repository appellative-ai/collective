package timeseries

import (
	http2 "github.com/behavioral-ai/core/http"
	"github.com/behavioral-ai/core/messaging"
)

var (
	Agent    messaging.Agent
	Exchange http2.Exchange
)

func init() {
	Exchange = http2.Do
}
