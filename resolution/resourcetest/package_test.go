package resourcetest

func ExampleNewResolver() {
	s := NewResolver()
	//m := messaging.NewAddressableMessage(messaging.ChannelControl, messaging.ConfigEvent, "core:to/test", "core:from/test")
	s.Representation("name")
	s.AddRepresentation("name", "author", "contentType", nil)
	s.Context("name")
	s.AddContext("name", "author", "contentType", nil)

	//Output:
	//message  -> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//advise   -> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//subscribe-> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//cancel   -> [chan:ctrl] [from:core:from/test] [to:core:to/test] [core:event/config]
	//2025-06-02T17:28:47.896Z [core:agent/operations/collective] [task] [going well] [none]

}
