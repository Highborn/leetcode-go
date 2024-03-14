package package202311

func computeMode(maxM, curM *[]int, val int) {
	if val == (*curM)[1] {
		(*curM)[0] += 1
	} else {
		if (*maxM)[0] < (*curM)[0] {
			*maxM = *curM
		} else if (*maxM)[0] == (*curM)[0] {
			*maxM = append(*maxM, (*curM)[1:]...)
		}
		*curM = []int{1, val}
	}
}

// FindMode suppose to find mode on a BST
func FindMode(root *TreeNode) []int {
	cNode := root
	maxM := []int{-1, 1e6}
	curM := []int{-1, 1e6}
	for cNode != nil {
		if cNode.Left != nil {
			newRoot := cNode.Left
			for newRoot.Right != nil && newRoot.Right != cNode {
				newRoot = newRoot.Right
			}
			newRoot.Right = cNode
			cNode = cNode.Left
			newRoot.Right.Left = nil
		} else {
			//			log.Println("L", cNode)
			computeMode(&maxM, &curM, cNode.Val)
			cNode = cNode.Right
		}
	}
	computeMode(&maxM, &curM, 1e6)
	return maxM[1:]
}

// MorrisTravers implements morris traverse of a tree and returns it in a slice
func MorrisTravers(root *TreeNode) []*TreeNode {
	curNode := root
	var traverse []*TreeNode
	for curNode != nil {
		if curNode.Left == nil {
			traverse = append(traverse, curNode)
			curNode = curNode.Right
		} else {
			newRoot := curNode.Left
			for newRoot.Right != nil && newRoot.Right != curNode {
				newRoot = newRoot.Right
			}
			if newRoot.Right == nil {
				newRoot.Right = curNode
				curNode = curNode.Left
			} else {
				newRoot.Right = nil
				traverse = append(traverse, curNode)
				curNode = curNode.Right
			}
		}
	}
	return traverse
}

func FindMode2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make(map[int]int, 2)
	traversal(root, result)
	out := make([]int, 0, len(result))
	max := 0
	for _, v := range result {
		if v > max {
			max = v
		}
	}
	for k, v := range result {
		if v >= max {
			out = append(out, k)
		}
	}
	return out
}

func traversal(root *TreeNode, counter map[int]int) {
	if root == nil {
		return
	}
	counter[root.Val] += 1
	traversal(root.Left, counter)
	traversal(root.Right, counter)
}
