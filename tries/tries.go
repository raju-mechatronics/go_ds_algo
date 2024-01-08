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
	childLen := len(t.children)
	if childLen == 1 && t.hasChild(rune(word[0])) {
		t.removeChild(rune(word[0]))
		return
	}

	if len(word) == 1 {
		t.children[rune(word[0])].isEnd = false
		return
	}

	t.children[rune(word[0])].Remove(word[1:])
}

func searching(n node, prevWord string) {

}

func (t *node) Search(word string) []string {
	cNode := t
	for i := 0; i < len(word); i++ {
		l := word[i]
		if cNode.hasChild(rune(l)) {
			cNode = cNode.children[rune(l)]
		} else {
			return make([]string, 0)
		}
	}

	return cNode.GetWords(word)
}

func (t *node) GetWords(prv string) []string {
	result := make([]string, 0)
	for n := range t.children {
		cNode := t.children[n]
		if cNode.isEnd {
			result = append(result, prv+string(cNode.value))
		}
		result = append(result, cNode.GetWords(prv+string(cNode.value))...)
	}
	return result
}
