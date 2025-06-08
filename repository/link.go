package repository

// newLinkMap - new exchange link map
func newLinkMap[T, U any]() *mapT[T, U] {
	return newMap[T, U]()
}
