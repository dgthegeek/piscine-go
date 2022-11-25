package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nod := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Right {
		nod.Parent.Right = rplc
	} else {
		nod.Parent.Left = rplc
	}
	nod.Parent = node.Parent
	return root
}
