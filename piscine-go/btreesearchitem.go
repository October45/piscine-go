package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else {
		return root
	}
}
