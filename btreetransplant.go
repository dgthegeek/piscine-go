package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rep := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		rep.Parent.Left = rplc
	} else {
		rep.Parent.Right = rplc
	}
	rep.Parent = node.Parent
	return root
}
