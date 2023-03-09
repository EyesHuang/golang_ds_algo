package tree

import "fmt"

type Node struct {
	value       int
	left, right *Node
}
type Tree struct {
	root *Node
}

// PrintTree
// Time Complexity: O(n)
func (t *Tree) PrintTree() {
	fmt.Print("Tree : ")
	printPreOrder(t.root)
	fmt.Println()
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)
}

// PrintPostOrder
// Time Complexity: O(n)
func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	}

	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Print(n.value, " ")
}

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}

	printInOrder(n.left)
	fmt.Print(n.value, " ")
	printInOrder(n.right)
}

// LevelOrderBinaryTree
// Time Complexity: O(n)
// Space Complexity: O(n)
func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}
