package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	} else if data > root.Data {
		return BTreeInsertData(root.Left, data)
	} else if data < root.Data {
		return BTreeInsertData(root.Right, data)
	}
	return root
}
