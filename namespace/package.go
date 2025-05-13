package namespace

// TODO : add activity
//        Can we add frame content version to activity?
//        Need author, frame, things, Accessor
//        Need counts, lists of aspects

// Accessor -
type Accessor struct {
	//Version string // returned on a Get
	Type    string // Content-Type
	Content any
}

// Adder - add
type Adder struct {
	Thing    func(name, cname, author string) *Status
	Relation func(name, cname, thing1, thing2, author, authority string) *Status
	// What exactly are the results?
	// How to query+select/return generational information
	// Content can be captured if provided.
	ConnectThing  func(name, frame, author string, access Accessor) (results string, status *Status)
	ConnectAspect func(name []string, frame, author string, access Accessor) (results string, status *Status)
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(name, cname, author string) *Status {
			return agent.addThing(name, cname, "", author)
		},
		Relation: func(name, cname, thing1, thing2, author, authority string) *Status {
			return agent.addRelation(name, cname, thing1, thing2, authority, author)
		},
	}
}()
