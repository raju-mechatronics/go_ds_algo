package hashmap

import (
	"ds_algo/linked_list"
	"fmt"
)

type iKey interface {
	string | int
}

type key_value[K iKey, V any] struct {
	key   K
	value V
}

type HashMap[K iKey, V any] struct {
	buckets [10]*linked_list.LinkedList[key_value[K, V]]
}

func hash(key interface{}) int {
	switch k := key.(type) {
	case string:
		hashCode := 0
		for _, char := range k {
			hashCode += int(char)
		}
		return hashCode % 10
	case int:
		return k % 10
	default:
		panic("Unsupported key type")
	}
}

func (m *HashMap[K, V]) Put(key K, value V) {
	hash_index := hash(key)
	if m.buckets[hash_index] == nil {
		new_list := linked_list.LinkedList[key_value[K, V]]{}
		new_list.AddLast(key_value[K, V]{key, value})
		m.buckets[hash_index] = &new_list
	} else {
		list := m.buckets[hash_index]
		next := list.Iterator()

		for cur := next(); cur != nil; cur = next() {
			if cur.key == key {
				cur.value = value
				return
			}
		}
		list.AddLast(key_value[K, V]{key, value})
	}
}

func (m *HashMap[K, V]) Get(key K) (V, error) {
	hash_index := hash(key)
	var zero V
	if m.buckets[hash_index] == nil {
		return zero, fmt.Errorf("key not found")
	}
	list := m.buckets[hash_index]
	next := list.Iterator()

	for cur := next(); cur != nil; cur = next() {
		if cur.key == key {
			return cur.value, nil
		}
	}
	return zero, fmt.Errorf("key not found")
}

func (m *HashMap[K, V]) Remove(key K) error {
	hash_index := hash(key)
	if m.buckets[hash_index] == nil {
		return fmt.Errorf("key not found")
	}
	list := m.buckets[hash_index]
	next := list.Iterator()

	i := 0
	for cur := next(); cur != nil; cur = next() {
		if cur.key == key {
			list.Remove(i)
			return nil
		}
		i++
	}
	return fmt.Errorf("key not found")
}
