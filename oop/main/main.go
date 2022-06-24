package main

import "github.com/huzhouv/go-learning/oop"

func main() {
	// m := new(M)
	// (&Sub{M: m.m1()}).s1()
	// oop.Run()

	oop.Trap()
}

type Sub struct {
	*M
	id string
}

func (s *Sub) s1() {
	println("haha: " + s.id)
}

type M struct {
	d *int
}

func (m *M) m1() *M {
	return m
}
