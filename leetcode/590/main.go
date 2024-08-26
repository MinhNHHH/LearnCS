package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// Input: root = [1,null,3,2,4,null,5,6]
// Output: [5,6,3,2,4,1]
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := []int{}
	var helper func(root *Node)
	helper = func(root *Node) {
		if root == nil {
			return
		}
		for _, neighboor := range root.Children {
			helper(neighboor)
		}
		ans = append(ans, root.Val)
	}
	helper(root)
	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
}
