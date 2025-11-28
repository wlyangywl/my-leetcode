package merge_trees_617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 合并根节点
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	curTreeNode := &TreeNode{
		Val: root1.Val + root2.Val,
	}
	// 合并左边子树
	curTreeNode.Left = mergeTrees(root1.Left, root2.Left)
	// 合并右边子树
	curTreeNode.Right = mergeTrees(root1.Right, root2.Right)
	return curTreeNode
}
