package namespace

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"sync/atomic"
)

// Namespace names format:
//  collective = domain-sub-domain
//  domain = domain-sub-domain
//
//  {collective}:{domain}:{type}/{path}#{fragment}
//
//  Example:  wikipedia-eng:resiliency-traffic:agent/rate-limiting/request/http#v1.2.3
//
// What would be a collective name in the root registry?
//     registry:{dns-name}:collective/{collective-name}
//
// TODO : add activity
//        Can we add frame resource version to activity?
//        Need author, frame, things, Accessor
//        Need counts, lists of aspects

const (
	Fragment = "#"
	Colon    = ":"
	Slash    = "/"
)

const (
	CollectiveKind = "collective" // References a collective in the registry/root collective
	ImportKind     = "import"     // References another collective to import
	HandlerKind    = "handler"    // Function used as a link in a Micro-REST chain
	LinkKind       = "link"       // Http URI

	AgentKind   = "agent"   // Used to define a thing that is empowered, agents are members of the collective, not just things.
	AspectKind  = "aspect"  // Used for making connections resources: self, info
	EventKind   = "event"   // Messaging events
	PersonKind  = "person"  // Used for authorization and ownership resources: self, info, instance
	ServiceKind = "service" // Service name
	TaskKind    = "task"    // Used for tracing agent activity. What is the agent tasked with
	TypeKind    = "type"    // Programming language types

	ThingKind    = "thing"    // Generic thing
	RelationKind = "relation" // Used for relating 2 resources
	FrameKind    = "frame"    // Used as a container for names, similar to a dir entry on a file system
)

var (
	counter = new(atomic.Int64)
)

// Name -
type Name struct {
	Collective string `json:"collective"`
	Domain     string `json:"domain"`
	Kind       string `json:"kind"`
	Path       string `json:"path"`
	Fragment   string `json:"fragment"`
}

func ParseName(name string) Name {
	return parse(name)
}

func AddFragment(name, fragment string) string {
	return addFragment(name, fragment)
}

func Versioned(name string) string {
	return AddFragment(name, fmt.Sprintf("%v", counter.Add(1)))
}

func Kind(name string) string {
	return ParseName(name).Kind
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
			return agent.addThing(name, cname, author)
		},
		Relation: func(name, cname, thing1, thing2, author string) *messaging.Status {
			return agent.addRelation(name, cname, thing1, thing2, author)
		},
	}
}()
