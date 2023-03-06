package main

import (
	"fmt"
	"runtime"
)

var (
	done chan struct{}
)

type A struct {
	name string
}

type B struct {
	name string
}

type C struct {
	name string
}

func newA() *A {
	v := &A{"n1"}
	runtime.SetFinalizer(v, func(p *A) {
		fmt.Println("gc Finalizer A")
	})
	return v
}

func newB() *B {
	v := &B{"n1"}
	runtime.SetFinalizer(v, func(p *B) {
		<-done
		fmt.Println("gc Finalizer B")
	})
	return v
}

func newC() *C {
	v := &C{"n1"}
	runtime.SetFinalizer(v, func(p *C) {
		fmt.Println("gc Finalizer C")
	})
	return v
}

func main() {
	a := newA()
	b := newB()
	c := newC()
	fmt.Println("== start ===")
	_, _, _ = a, b, c
	fmt.Println("== ... ===")
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("== end ===")
}
