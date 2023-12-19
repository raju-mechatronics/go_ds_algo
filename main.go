package main

import "fmt"

type Fest struct {
	n int
}

type Test struct {
	f *Fest
}

func F2(i *Fest) {
	fmt.Printf(", %p \n", i)
	i = &Fest{0}
	fmt.Printf(", %p \n", i)
}

func F1(t *Test) {
	fmt.Printf("%p %p", t, t.f)
	F2(t.f)
}

func main() {
	t := &Test{&Fest{1}}
	fmt.Printf("%p \n", t)
	F1(t)
	fmt.Println(t.f.n)
}
