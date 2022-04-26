package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}
	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}
	successor := BTreeMin(node.Right)
	if successor.Parent != node {
		root = BTreeTransplant(root, successor, successor.Right)
		successor.Right = node.Right
		successor.Right.Parent = successor
	}
	return BTreeTransplant(root, node, successor)
}
