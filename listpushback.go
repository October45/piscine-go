package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, n interface{}) {
	node := &NodeL{Data: n}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
}
