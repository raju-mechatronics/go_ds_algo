package tries

import "fmt"

type node struct {
	value    rune
	children map[rune]*node
	isEnd    bool
}

func (n *node) hasChild(r rune) bool {
	_, ok := n.children[r]
	return ok
}

func (n *node) insertChild(r rune) {
	n.children[r] = &node{value: r, children: make(map[rune]*node)}
}

func (n node) String() string {
	children := ""
	for _, child := range n.children {
		children += child.String() + " "
	}
	return fmt.Sprintf("%c [%s]", n.value, children)
}

func New() node {
	return node{value: 0, children: make(map[rune]*node), isEnd: false}
}

func (t *node) Insert(s string) {
	if len(s) == 0 {
		return
	}
	if len(s) == 1 {
		r := rune(s[0])
		if t.hasChild(r) {
			t.children[r].isEnd = true
		} else {
			t.children[r] = &node{
				value:    r,
				children: make(map[rune]*node),
				isEnd:    true,
			}
		}
		return
	}
	r := rune(s[0])

	if !t.hasChild(r) {
		t.insertChild(r)
	}
	t.children[r].Insert(s[1:])
}

func (t *node) removeChild(r rune) {
	delete(t.children, r)
}

func (t *node) Remove(word string) {

}
