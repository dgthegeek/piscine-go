package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data > elem {
		return BTreeSearchItem(root.Right, elem)
	} else if root.Data < elem {
		return BTreeSearchItem(root.Left, elem)
	}
	return root
}
