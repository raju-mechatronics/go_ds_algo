package main

import (
	"ds_algo/hashmap"
	"fmt"
)

func main() {
	//check hash map
	m := hashmap.HashMap[string, int]{}
	m.Put("hello", 1)
	m.Put("world", 2)
	m.Put("hello", 3)
	// check if put is working
	fmt.Println(m.Get("hello"))
	fmt.Println(m.Get("world"))
	// check if get is working
	fmt.Println(m.Get("not_found"))
}
