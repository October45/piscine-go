package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	size := 1
	for l.Head.Next != nil {
		size++
		l.Head = l.Head.Next
	}
	return size
}
