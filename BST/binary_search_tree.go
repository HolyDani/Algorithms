package BST

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert new value in tree
func (n *TreeNode) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search value in tree
func (n *TreeNode) Search(value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// Recursive tree traversal in LNR order
// Good for sorted value from tree
func (n *TreeNode) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrderTraversal()
}
