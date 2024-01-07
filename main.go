package main

import "ds_algo/tries"

func main() {
	s := tries.New()
	s.Insert("hello")
	s.Insert("hallo")
	s.Insert("hellow")
}
