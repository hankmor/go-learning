package main

import "fmt"

func main() {
	var p ids
	// always copy it when exec each method
	ret := p.add(1, 2, 3).f1()
	fmt.Printf("%d, %d\n", &p, &ret) // &[], &[1 2 3]
	pp := &idsp{}
	// address of pointer
	ret2 := pp.add(1, 2, 3)
	// address of pointer
	ret3 := ret2.f1()
	fmt.Printf("%d, %d, %d\n", &pp, &ret2, &ret3) // 824634417184, 824634417192
	fmt.Println(pp, ret2, ret3)                   // &[1 2 3] &[1 2 3] &[1 2 3]
}

type ids []int

func (i ids) add(is ...int) ids {
	i = append(i, is...)
	return i
}

func (i ids) f1() ids {
	// do something
	return i
}

type idsp []int

func (i *idsp) add(is ...int) *idsp {
	*i = append(*i, is...) // copy
	return i
}

func (i *idsp) f1() *idsp {
	// do something
	return i
}
