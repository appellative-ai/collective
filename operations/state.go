package operations

import "github.com/behavioral-ai/core/messaging"

type operationsT struct {
	running           bool
	primaryHost       string
	secondaryHost     string
	serviceName       string
	collectiveName    string
	linkedCollectives map[string][]string
	origin            originT
}

func initialize(msg *messaging.Message) (ops *operationsT, ok bool) {
	ops = new(operationsT)
	if ops.origin, ok = newOriginFromMessage(msg); !ok {
		return
	}
	ops.serviceName = ops.origin.Name()
	ops.collectiveName = ops.origin.AppId
	cfg := messaging.ConfigMapContent(msg)
	if cfg != nil {
		ops.primaryHost = cfg[PrimaryHost]
		ops.secondaryHost = cfg[SecondaryHost]
	}
	return
}
