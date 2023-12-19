package binary_tree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	BT := BinaryTree{nil}
	BT.Insert(5)
	fmt.Println(BT)
	if BT.root.value != 5 {
		t.Errorf("Insert failed, expected %v, got %v", 5, BT.root.value)
	}
	BT.Insert(3)
	if BT.root.leftChild.value != 3 {
		t.Errorf("Insert failed, expected %v, got %v", 3, BT.root.leftChild.value)
	}
	BT.Insert(7)
	if BT.root.rightChild.value != 7 {
		t.Errorf("Insert failed, expected %v, got %v", 7, BT.root.rightChild.value)
	}
	BT.Insert(4)
	if BT.root.leftChild.rightChild.value != 4 {
		t.Errorf("Insert failed, expected %v, got %v", 4, BT.root.leftChild.rightChild.value)
	}
	BT.Insert(6)
	if BT.root.rightChild.leftChild.value != 6 {
		t.Errorf("Insert failed, expected %v, got %v", 6, BT.root.rightChild.leftChild.value)
	}
}

func TestGetHeight(t *testing.T) {
	BT := BinaryTree{nil}
	BT.Insert(5)
	BT.Insert(3)
	BT.Insert(7)
	BT.Insert(4)
	BT.Insert(6)
	if BT.GetHeight() != 3 {
		t.Errorf("GetHeight failed, expected %v, got %v", 3, BT.GetHeight())
	}
	BT.Insert(10)
	BT.Insert(11)
	if BT.GetHeight() != 4 {
		t.Errorf("GetHeight failed, expected %v, got %v", 4, BT.GetHeight())
	}
}

func TestGetMax(t *testing.T) {
	BT := BinaryTree{nil}
	BT.Insert(5)
	BT.Insert(3)
	BT.Insert(7)
	BT.Insert(4)
	BT.Insert(6)
	if maxNumber, err := BT.GetMax(); maxNumber != 7 || err != nil {
		t.Errorf("GetMax failed, expected %v, got %v", 7, maxNumber)
	}
}

func TestBinaryTree_GetMin(t *testing.T) {
	BT := BinaryTree{nil}
	BT.Insert(5)
	BT.Insert(3)
	BT.Insert(7)
	BT.Insert(4)
	BT.Insert(6)
	if minNumber, err := BT.GetMin(); minNumber != 3 || err != nil {
		t.Errorf("GetMax failed, expected %v, got %v", 3, minNumber)
	}
}

// test contains
func TestBinaryTree_Contains(t *testing.T) {
	BT := BinaryTree{nil}
	BT.Insert(5)
	BT.Insert(3)
	BT.Insert(7)
	BT.Insert(4)
	BT.Insert(6)
	if !BT.Contains(3) {
		t.Errorf("Contains failed, expected %v, got %v", true, BT.Contains(3))
	}
	if BT.Contains(10) {
		t.Errorf("Contains failed, expected %v, got %v", false, BT.Contains(10))
	}
}

// test: get ancestor. a value is given, find the ancestor of this value. GetAncestors returns the ancestor values in a slice.
func TestBinaryTree_GetAncestor(t *testing.T) {
	BT := BinaryTree{nil}
	BT.Insert(5)
	BT.Insert(3)
	BT.Insert(7)
	BT.Insert(4)
	BT.Insert(6)
	BT.Insert(9)
	BT.Insert(10)
	BT.Insert(8)
	ancestors := BT.GetAncestors(8)
	if ancestors[0] != 9 || ancestors[1] != 7 || ancestors[2] != 5 {
		t.Errorf("GetAncestors failed, expected %v, got %v", []int{9, 7, 5}, ancestors)
	}
}

/*
		5
	3		   7
		4	6	  9
				8	10

*/
