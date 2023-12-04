package hashmap

import "testing"

func TestHashMap(t *testing.T) {
	m := HashMap[string, int]{}
	m.Put("hello", 1)
	m.Put("world", 2)
	m.Put("hello", 3)
	// check if put is working
	if v, _ := m.Get("hello"); v != 3 {
		t.Errorf("Expected 3, got %v", v)
	}
	if v, _ := m.Get("world"); v != 2 {
		t.Errorf("Expected 2, got %v", v)
	}
	// check if get is working
	if _, err := m.Get("not_found"); err == nil {
		t.Errorf("Expected error, got nil")
	}
}
