package namespace

type relation struct {
	name   string
	thing1 string
	thing2 string
}

func (r relation) contains(thing string) bool { return r.thing1 == thing || r.thing2 == thing }

type relationT struct {
	items []relation
}

func newRelation() *relationT {
	c := new(relationT)
	return c
}

func (r *relationT) find(thing string) (relation, bool) {
	for i, item := range r.items {
		if item.contains(thing) {
			return r.items[i], true
		}
	}
	return relation{}, false
}

func (r *relationT) put(name, thing1, thing2 string) {
	r.items = append(r.items, relation{name: name, thing1: thing1, thing2: thing2})
}
