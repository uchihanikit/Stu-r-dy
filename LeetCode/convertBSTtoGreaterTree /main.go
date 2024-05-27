package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

var sum = 0

func (t *Tree) Insert(value int) {
	newNode := &TreeNode{Val: value}

	if t.root == nil {
		t.root = newNode
		return
	}

	current := t.root
	parent := t.root

	for {
		parent = current
		if value < current.Val {
			current = current.Left
			if current == nil {
				parent.Left = newNode
				return
			}
		} else {
			current = current.Right
			if current == nil {
				parent.Right = newNode
				return
			}
		}
	}
}

func PrintTree(root *TreeNode) {
	fmt.Print(root.Val, " ")
	if root.Left != nil {
		PrintTree(root.Left)
	}
	if root.Right != nil {

		PrintTree(root.Right)
	}
	return
}

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	convertBSTHelper(root)
	return root
}

func convertBSTHelper(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTHelper(root.Right)
	sum += root.Val
	root.Val = sum
	convertBSTHelper(root.Left)
}

func main() {
	input := []int{4, 1, 6, 0, 2, 5, 7, 3, 8}
	tree := &Tree{}
	for _, val := range input {
		tree.Insert(val)
	}
	PrintTree(tree.root)
	fmt.Println()
	answer := convertBST(tree.root)
	PrintTree(answer)
	input = []int{0, 1}
	tree = &Tree{}
	for _, val := range input {
		tree.Insert(val)
	}
	PrintTree(tree.root)
	fmt.Println()
	answer = convertBST(tree.root)
	PrintTree(answer)
}
