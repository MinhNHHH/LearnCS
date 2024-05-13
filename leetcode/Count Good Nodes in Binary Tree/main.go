// Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

// Return the number of good nodes in the binary tree.

// Input: root = [3,1,4,3,null,1,5]
// Output: 4
// Explanation: Nodes in blue are good.
// Root Node (3) is always a good node.
// Node 4 -> (3,4) is the maximum value in the path starting from the root.
// Node 5 -> (3,4,5) is the maximum value in the path
// Node 3 -> (3,1,3) is the maximum value in the path.

// Input: root = [3,3,null,4,2]
// Output: 3
// Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(root *TreeNode, maxVal int) int {
	res := 0

	if root == nil {
		return 0
	}

	if root.Val >= maxVal {
		res = 1
	}
	maxVal = max(maxVal, root.Val)
	res += dfs(root.Left, maxVal)
	res += dfs(root.Right, maxVal)
	return res
}

func goodNodes(root *TreeNode) int {

	return dfs(root, root.Val)
}
