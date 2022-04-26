package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head.Next != nil {
		l.Head = l.Head.Next
	}
	return l.Head.Data
}
