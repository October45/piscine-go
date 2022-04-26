package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}
	if l.Data > node.Data {
		node.Next = l
		return node
	}
	iterator := l
	for iterator.Next != nil && iterator.Next.Data < node.Data {
		iterator = iterator.Next
	}
	node.Next = iterator.Next
	iterator.Next = node
	return l
}
