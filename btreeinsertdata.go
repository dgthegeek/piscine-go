package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right.Parent, data)
	} else if data < root.Data {
		root.Left = BTreeInsertData(root.Left.Parent, data)
	}
	return root
}
