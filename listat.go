package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	for i := 0; i < pos; i++ {
		if l.Next == nil {
			return nil
		}
		l = l.Next
	}
	return l
}
