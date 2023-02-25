package tree_test

import (
	algoTree "golang_ds_algo/tree"
	"testing"
)

func Test_Tree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := algoTree.LevelOrderBinaryTree(arr)
	tree.PrintTree()
}
