package main

import (
	"ds_algo/linked_list"
	"fmt"
)

func main() {
	linkedList := linked_list.LinkedList{}
	linkedList.AddLast(1)
	linkedList.AddLast(2)
	linkedList.AddLast(3)
	linkedList.AddLast(4)
	linkedList.AddLast(5)
	linkedList.AddLast(6)
	linkedList.AddLast(7)
	linkedList.AddLast(8)
	linkedList.AddLast(9)
	linkedList.AddLast(1)
	fmt.Println(linkedList.ToArray())
	linkedList.Reverse()
	fmt.Println(linkedList.ToArray())
}
