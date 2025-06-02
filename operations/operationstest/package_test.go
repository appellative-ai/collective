package operationstest

import (
	"github.com/behavioral-ai/core/messaging"
)

func ExampleNewService() {
	s := NewService()
	m := messaging.NewAddressableMessage(messaging.ChannelControl, messaging.ConfigEvent, "core:to/test", "core:from/test")
	s.Message(m)
	s.Advise(m)
	s.Subscribe(m)
	s.CancelSubscription(m)
	s.Trace("core:agent/operations/collective", "task", "going well", "none")

	//Output:
	//message  -> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//advise   -> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//subscribe-> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//cancel   -> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//2025-06-02T17:28:47.896Z [core:agent/operations/collective] [task] [going well] [none]

}
