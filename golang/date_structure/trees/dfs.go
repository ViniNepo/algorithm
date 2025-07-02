package trees

import "fmt"

type DFS struct{}

func (d *DFS) InOrder(root *TreeNode) {
	if root == nil {
		return
	}

	d.InOrder(root.Left)
	fmt.Println(root.Val)
	d.InOrder(root.Right)
}

func (d *DFS) PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	d.PreOrder(root.Left)
	d.PreOrder(root.Right)
}

func (d *DFS) PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	d.PostOrder(root.Left)
	d.PostOrder(root.Right)
	fmt.Println(root.Val)
}
