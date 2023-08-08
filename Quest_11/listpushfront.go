package piscine

/*
Write a function ListPushFront that inserts a new element NodeL at the beginning of the list l while using the structure List
*/

type aNodeL struct {
	Data interface{}
	Next *NodeL
}

type aList struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	uus_node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = uus_node
		return
	}

	uus_node.Next = l.Head
	l.Head = uus_node
}
