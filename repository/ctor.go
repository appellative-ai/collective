package repository

// ctorM - constructor map
//type ctorM struct {
//	m *sync.Map
//}

// newCtorMap - new agent func map
func newCtorMap[T, U any]() *mapT[T, U] {
	return newMap[T, U]()
}

/*
func (c *ctorM) get(name string) messaging.NewAgent {
	v, ok := c.m.Load(name)
	if !ok {
		return nil
	}
	if v1, ok1 := v.(messaging.NewAgent); ok1 {
		return v1
	}
	return nil
}

func (c *ctorM) store(name string, fn messaging.NewAgent) {
	if name == "" || fn == nil {
		return
	}
	c.m.Store(name, fn)
}


*/
