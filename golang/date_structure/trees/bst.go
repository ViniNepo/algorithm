package trees

type BST struct{}

func (b *BST) Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}

	if val > root.Val {
		root.Right = b.Insert(root.Right, val)
	} else if val < root.Val {
		root.Left = b.Insert(root.Left, val)
	}

	return root
}

func (b *BST) MinValue(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func (b *BST) Remove(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		root.Right = b.Remove(root.Right, val)
	} else if val < root.Val {
		root.Left = b.Remove(root.Left, val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := b.MinValue(root.Right)
			root.Val = minNode.Val
			root.Right = b.Remove(root.Right, minNode.Val)
		}
	}

	return root
}
