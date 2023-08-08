package piscine

// finds the element and returns the first data from the node that is a string
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	link := l.Head
	for link != nil {
		if comp(link.Data, ref) {
			return &link.Data
		}

		link = link.Next
	}
	return nil
}

// defines for two elements the equality criteria
func CompStr(a, b interface{}) bool {
	return a == b
}
