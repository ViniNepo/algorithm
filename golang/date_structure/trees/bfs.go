package trees

import (
	"fmt"
	"github.com/gammazero/deque"
)

type BFS struct{}

func NewBFS() *BFS {
	return &BFS{}
}

func (b *BFS) bfs(root *TreeNode) {
	var queue deque.Deque[*TreeNode]
	if root != nil {
		queue.PushBack(root)
	}

	level := 0
	for queue.Len() != 0 {
		fmt.Printf("level %d: ", level)
		levelLength := queue.Len()

		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront()
			fmt.Printf("%d", curr.Val)

			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}

			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}

		level++
		fmt.Println()
	}
}
