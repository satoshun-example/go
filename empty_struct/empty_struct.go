package main

import (
	"fmt"
	"unsafe"
)

type empty struct {
	e struct{}
}

func (e *empty) addr() {
	fmt.Printf("addr:%p\n", e)
}

func main() {
	var e1, e2 empty

	fmt.Println(e1)
	fmt.Println(unsafe.Sizeof(e1))

	e1.addr()
	e2.addr()
	fmt.Println(&e1 == &e2) // it's true!
}
