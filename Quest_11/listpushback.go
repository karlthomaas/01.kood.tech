package piscine

/*
Write a function ListPushBack that inserts a new element NodeL at the end of the list l
while using the structure List.
*/

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	uus_node := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = uus_node
	} else {
		l.Tail.Next = uus_node
	}
	l.Tail = uus_node
}
