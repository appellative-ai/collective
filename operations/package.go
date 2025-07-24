package operations

import "github.com/appellative-ai/core/messaging"

const (
// ServiceKind = "service"
)

// Origin map and host keys
const (
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
)

// Service - servicing add functions
type Service struct {
	AddRepresentation func(name, author, contentType string, value any) *messaging.Status
	AddContext        func(name, author, contentType string, value any) *messaging.Status

	AddThing func(name, cname, author string) *messaging.Status
	AddJoin  func(name, cname, thing1, thing2, author string) *messaging.Status
}

// Serve -
var Serve = func() *Service {
	return &Service{
		AddRepresentation: func(name, author, contentType string, value any) *messaging.Status {
			return messaging.StatusOK()
		},
		AddContext: func(name, author, contentType string, t any) *messaging.Status {
			return messaging.StatusOK()
		},
		AddThing: func(name, cname, author string) *messaging.Status {
			return messaging.StatusOK()
		},
		AddJoin: func(name, cname, thing1, thing2, author string) *messaging.Status {
			return messaging.StatusOK()
		},
		/*
			AddOrderJoin: func(name, cname, thing1, thing2, author string) *messaging.Status {
				return messaging.StatusOK()
			},

		*/
		/*
			SubscriptionCreate: func(msg *messaging.Message) {
				agent.subscribe(msg)
			},
			SubscriptionCancel: func(msg *messaging.Message) {
				agent.cancel(msg)
			},

		*/
	}
}()

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}

// Accessor -
type Accessor struct {
	//Version string // returned on a Get
	Type    string // Content-Type
	Content any
}

// Adder - add
type Adder struct {
	Thing    func(name, cname, author string) *messaging.Status
	Relation func(name, cname, thing1, thing2, author string) *messaging.Status
	// What exactly are the results?
	// How to query+select/return generational information
	// Content can be captured if provided.
	ConnectThing  func(name, frame, author string, access Accessor) (results string, status *messaging.Status)
	ConnectAspect func(name []string, frame, author string, access Accessor) (results string, status *messaging.Status)
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(name, cname, author string) *messaging.Status {
			return messaging.StatusOK() //agent.addThing(name, cname, author)
		},
		Relation: func(name, cname, thing1, thing2, author string) *messaging.Status {
			return messaging.StatusOK() //agent.addRelation(name, cname, thing1, thing2, author)
		},
	}
}()
