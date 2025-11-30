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
	var maxDiameter int
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 左子树的最大深度
		leftDepth := dfs(node.Left)
		// 右子树的最大深度
		rightDepth := dfs(node.Right)
		// 当前节点的最大直径和所有节点的最大直径对比一下
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)
		return max(leftDepth+rightDepth) + 1
	}
	dfs(root)
	return maxDiameter
}
