package diameter_of_binaryTree_543

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。
*/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDept := maxDepthTree(root.Left)
	rightDept := maxDepthTree(root.Right)

	maxDiameter := diameterOfBinaryTree(root.Left) + diameterOfBinaryTree(root.Right)
	return max(leftDept+rightDept, maxDiameter)
}

func maxDepthTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepthTree(root.Left)
	rightDepth := maxDepthTree(root.Right)

	if rightDepth > leftDepth {
		return rightDepth + 1
	}
	return leftDepth + 1
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
