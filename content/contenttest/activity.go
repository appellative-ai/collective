package contenttest

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

func AddActivity(hostName string, agent messaging.Agent, event, source string, content any) {
	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
}
