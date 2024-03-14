package package202311

import (
	"testing"
)

var testCases = []struct {
	treeArray []int
	expected  int
}{
	{[]int{4, 8, 5, 0, 1, -1, 6}, 5},
	{[]int{1}, 1},
}

func createTree(arr []int) (root *TreeNode) {
	root = &TreeNode{arr[0], nil, nil}
	nodeList := make([]*TreeNode, len(arr))
	nodeList[0] = root
	for idx, value := range arr[1:] {
		if value < 0 {
			continue
		}
		node := &TreeNode{
			value,
			nil,
			nil,
		}
		nodeList[idx+1] = node
		parent := nodeList[idx/2]
		if idx%2 == 0 {
			parent.Left = node
		} else {
			parent.Right = node
		}
	}
	return root
}

func TestAverageOfSubtree(t *testing.T) {
	for _, tt := range testCases {
		root := createTree(tt.treeArray)
		got := AverageOfSubtree(root)
		if tt.expected != got {
			t.Error("Does not fit. Expected:\t", tt.expected, "\t,got:\t", got)
		}
	}
}
