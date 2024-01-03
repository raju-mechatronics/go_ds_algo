package heap

import (
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
	rIndex := h.getRightChildIndex(index)
	if rIndex >= 0 {
		return h.arr[rIndex]
	} else {
		return -1
	}
}

func (h *Heap) getLeftChild(index int) int {
	rIndex := h.getLeftChildIndex(index)
	if rIndex >= 0 {
		return h.arr[rIndex]
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

func (h *Heap) getParent(index int) int {
	pIndex := h.getParentIndex(index)
	if pIndex > -1 {
		return h.arr[pIndex]
	} else {
		return -1
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
	h.BubbleUp()
}

func (h *Heap) Remove() int {
	if h.size == 0 {
		log.Fatalf("Heap is empty")
	}
	start, end := h.arr[0], h.arr[h.size-1]
	h.arr[0], h.arr[h.size-1] = end, start
	h.size = h.size - 1
	h.arr = h.arr[:h.size]
	h.BubbleDown()
	return start
}

func (h Heap) Sort() []int {
	var sorted []int
	for i := 0; i < h.size; i++ {
		sorted = append(sorted, h.Remove())
	}
	return sorted
}

func (h *Heap) BubbleUp() {
	size := h.size
	for i := size - 1; i >= 0; {
		if h.hasParent(i) && h.get(i) > h.getParent(i) {
			h.arr[i], h.arr[h.getParentIndex(i)] = h.arr[h.getParentIndex(i)], h.arr[i]
			i = h.getParentIndex(i)
		} else {
			break
		}
	}
}

func (h *Heap) hasParent(i int) bool {
	if i > 0 {
		return true
	} else {
		return false
	}
}

func (h *Heap) BubbleDown() {
	size := h.size

	for i := 0; i < size; {
		if h.hasChildreen(i) {
			if h.hasRightChild(i) {
				if h.get(i) >= h.getLeftChild(i) && h.get(i) >= h.getRightChild(i) {
					break
				} else if h.getLeftChild(i) > h.getRightChild(i) {
					h.arr[i], h.arr[h.getLeftChildIndex(i)] = h.arr[h.getLeftChildIndex(i)], h.arr[i]
					i = h.getLeftChildIndex(i)
					continue
				} else {
					h.arr[i], h.arr[h.getRightChildIndex(i)] = h.arr[h.getRightChildIndex(i)], h.arr[i]
					i = h.getRightChildIndex(i)
					continue
				}
			} else {
				if h.get(i) < h.getLeftChild(i) {
					h.arr[i], h.arr[h.getLeftChildIndex(i)] = h.arr[h.getLeftChildIndex(i)], h.arr[i]
					i = h.getLeftChildIndex(i)
					continue
				} else {
					break
				}
			}
		}
		break
	}
}
