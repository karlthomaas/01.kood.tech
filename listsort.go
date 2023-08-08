package piscine

func ListSort(l *NodeI) *NodeI {
	l
	if l == nil {
		return nil
	}
	l.Next = ListSort(l.Next)

	if l.Next != nil && l.Data > l.Next.Data {
		l = move(l)
	}
	return l
}

func move(l *NodeI) *NodeI {
	u := l
	n := l.Next
	result := n

	for n != nil && l.Data > n.Data {
		u = n
		n = n.Next
	}
	u.Next = l
	l.Next = n
	return result
}
