package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	var prev *NodeL
	for l.Head != nil {
		l.Head.Next, prev, l.Head = prev, l.Head, l.Head.Next
	}
	l.Head = prev
}
