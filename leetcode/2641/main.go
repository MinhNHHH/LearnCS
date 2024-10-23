package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type Pair struct {
	Level int
	Node  *TreeNode
	Sum   int
}

func generateBinaryTree(temp []interface{}, root *TreeNode, i int) *TreeNode {
	if i < len(temp) && temp[i] != nil {
		node := &TreeNode{Val: temp[i].(int)}
		root = node

		root.Left = generateBinaryTree(temp, root.Left, i*2+1)

		root.Right = generateBinaryTree(temp, root.Right, i*2+2)
	}
	return root
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	temp := root
	stack := []*TreeNode{temp}
	level := 0
	trackingPair := []*Pair{{level, root, root.Val}}
	levelSums := []int{root.Val}

	for len(stack) > 0 {
		stackSize := len(stack)
		levelSum := 0
		for i := 0; i < stackSize; i++ {
			curr := stack[0]
			stack = stack[1:]
			slibSum := 0

			if curr.Left != nil {
				slibSum += curr.Left.Val
			}
			if curr.Right != nil {
				slibSum += curr.Right.Val
			}

			if curr.Left != nil {
				stack = append(stack, curr.Left)
				trackingPair = append(trackingPair, &Pair{level + 1, curr.Left, slibSum})
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
				trackingPair = append(trackingPair, &Pair{level + 1, curr.Right, slibSum})
			}
			levelSum += slibSum
		}
		levelSums = append(levelSums, levelSum)
		level++
	}
	for len(trackingPair) > 0 {
		curr := trackingPair[0]
		trackingPair = trackingPair[1:]
		level := curr.Level
		node := curr.Node
		sum := curr.Sum
		node.Val = levelSums[level] - sum
	}
	return root
}
func main() {
	arr := []interface{}{5, 4, 9, 1, 10, nil, 7}
	var root *TreeNode
	root = generateBinaryTree(arr, root, 0)
	replaceValueInTree(root)
}
