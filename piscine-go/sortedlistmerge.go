package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data < n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	}
	n2.Next = SortedListMerge(n1, n2.Next)
	return n2
}
