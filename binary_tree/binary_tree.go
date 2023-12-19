package binary_tree

import "errors"

type node struct {
	value      int
	leftChild  *node
	rightChild *node
}

func createNewNode(value int) *node {
	return &node{value, nil, nil}
}

type BinaryTree struct {
	root *node
}

func insert(root *node, value int) *node {
	if root == nil {
		root = createNewNode(value)
		return root
	}
	if value < root.value {
		root.leftChild = insert(root.leftChild, value)
	} else {
		root.rightChild = insert(root.rightChild, value)
	}
	return root
}

func (BT *BinaryTree) Insert(value int) {
	BT.root = insert(BT.root, value)
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
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	}
	if root.value > value {
		return contains(root.leftChild, value)
	} else {
		return contains(root.rightChild, value)
	}
}

func (BT BinaryTree) Contains(value int) bool {
	return contains(BT.root, value)
}

func getAncestor(root *node, value int) []int {
	if root == nil {
		return []int{}
	}
	if root.value == value {
		return []int{}
	}
	if root.leftChild != nil && root.leftChild.value == value {
		return []int{root.value}
	} else if root.rightChild != nil && root.rightChild.value == value {
		return []int{root.value}
	} else {
		leftAncestors := getAncestor(root.leftChild, value)
		if len(leftAncestors) != 0 {
			return append(leftAncestors, root.value)
		}
		rightAncestors := getAncestor(root.rightChild, value)
		if len(rightAncestors) != 0 {
			return append(rightAncestors, root.value)
		}
		return []int{}
	}
}

func (BT BinaryTree) GetAncestors(value int) []int {
	return getAncestor(BT.root, value)
}
