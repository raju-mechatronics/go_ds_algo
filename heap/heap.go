package heap

import (
	"fmt"
	"log"
	"net"
)

type Heap struct {
	arr  []int
	size int
}

func (h *Heap) GetArr() []int {
	return h.arr
}

func New() Heap {
	return Heap{make([]int, 0), 0}
}

func (h *Heap) getLeftChildIndex(index int) int {
	if index*2 < h.size {
		return index * 2
	} else {
		return -1
	}
}

func (h *Heap) getRightChildIndex(index int) int {

	if index*2+1 < h.size {
		return index*2 + 1
	} else {
		return -1
	}
}

func (h *Heap) getRightChild(index int) int {
	r_index := h.getRightChildIndex(index)
	if r_index >= 0 {
		return h.arr[r_index]
	} else {
		return -1
	}
}

func (h *Heap) getLeftChild(index int) int {
	r_index := h.getLeftChildIndex(index)
	if r_index >= 0 {
		return h.arr[r_index]
	} else {
		return -1
	}
}

func (h *Heap) getParentIndex(index int) int {
	if index > 0 {
		return index / 2
	} else {
		return -1
	}
}

func (h *Heap) getParent(index int) (int, error) {
	p_index := h.getParentIndex(index)
	if p_index > -1 {
		return h.arr[p_index], nil
	} else {
		return -1, fmt.Errorf("this is the root")
	}
}

func (h *Heap) hasChildreen(index int) bool {
	if ind := h.getLeftChildIndex(index); ind == -1 {
		return false
	} else {
		return true
	}
}

func (h *Heap) hasRightChild(index int) bool {
	if ind := h.getRightChildIndex(index); ind == -1 {
		return false
	} else {
		return true
	}
}

func (h *Heap) get(index int) int {
	return h.arr[index]
}

func (h *Heap) rearrange() interface{} {
	for i := 0; i < h.size; i++ {
		if h.hasChildreen(i) {
			if h.hasRightChild(i) {
				if h.get(i) >= h.getLeftChild(i) && h.get(i) >= h.getRightChild(i) {
					continue
				} else if h.getLeftChild(i) > h.getRightChild(i) {
					h.arr[i], h.arr[h.getLeftChildIndex(i)] = h.arr[h.getLeftChildIndex(i)], h.arr[i]
				} else {
					h.arr[i], h.arr[h.getRightChildIndex(i)] = h.arr[h.getRightChildIndex(i)], h.arr[i]
				}
			} else {
				if h.get(i) < h.getLeftChild(i) {
					h.arr[i], h.arr[h.getLeftChildIndex(i)] = h.arr[h.getLeftChildIndex(i)], h.arr[i]
				} else {
					continue
				}
			}
			h.rearrange()
		}
	}
	return net.Interface{}
}

func (h *Heap) Insert(i int) {
	h.size = h.size + 1
	h.arr = append(h.arr, i)
	h.rearrange()
}

func (h *Heap) Remove() int {
	if h.size == 0 {
		log.Fatalf("Heap is empty")
	}
	x := h.arr[0]
	h.arr = h.arr[1:]
	h.size = h.size - 1
	h.rearrange()
	return x
}

func (h Heap) Sort() []int {
	var sorted []int
	for i := 0; i < h.size; i++ {
		sorted = append(sorted, h.Remove())
	}
	return sorted
}

//
//func (h *Heap) Pop() int {
//
//}
//
//func (h *Heap) Sort() {
//
//}
