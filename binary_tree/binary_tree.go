package binary_tree

import "errors"

type node struct {
	value      int
	leftChild  *node
	rightChild *node
}

type BinaryTree struct {
	root *node
}

func insert(root *node, value int) {
	if root == nil {
		newNode := node{value: value}
		*root = newNode
		return
	}
	if value < root.value {
		insert(root.leftChild, value)
	} else {
		insert(root.rightChild, value)
	}
}

func (BT *BinaryTree) Insert(value int) {
	insert(BT.root, value)
}

func getHeight(root *node) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.leftChild), getHeight(root.rightChild)) + 1
}
func (BT BinaryTree) GetHeight() int {
	return getHeight(BT.root)
}

func getMax(root *node) (int, error) {
	if root == nil {
		return 0, errors.New("tree is empty")
	} else {
		if root.rightChild == nil {
			return root.value, nil
		} else {
			return getMax(root.rightChild)
		}
	}
}

func (BT BinaryTree) GetMax() (int, error) {
	return getMax(BT.root)
}

func getMin(root *node) (int, error) {
	if root == nil {
		return 0, errors.New("tree is empty")
	} else {
		if root.leftChild == nil {
			return root.value, nil
		} else {
			return getMin(root.leftChild)
		}
	}
}

func (BT BinaryTree) GetMin() (int, error) {
	return getMin(BT.root)
}

func countLeaves(root *node) int {
	if root == nil {
		return 0
	}
	if root.leftChild == nil && root.rightChild == nil {
		return 1
	} else {
		return countLeaves(root.leftChild) + countLeaves(root.rightChild)
	}
}

func (BT BinaryTree) CountLeaves() int {
	return countLeaves(BT.root)
}

func contains(root *node, value int) bool {
	if root.value == value {
		return true
	}
	if root == nil {
		return false
	}
	return contains(root.leftChild, value) || contains(root.rightChild, value)
}

func (BT BinaryTree) Contains(value int) bool {
	return contains(BT.root, value)
}

func getAncestor(root *node, value int) bool {
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	}
	return getAncestor(root.leftChild, value) || getAncestor(root.rightChild, value)
}

func (BT BinaryTree) GetAncestors() []int {

}
