package main

import "fmt"

type shape interface {
	sum() int
}

type squ struct {
	a int
}

type input struct {
	a int
	b int
}

func main() {
	var i shape
	i = &input{a: 10, b: 20}
	fmt.Println(i.sum())
	i = &squ{a: 10}
	fmt.Println(i.sum())
}

func (i *input) sum() int {
	return i.a + i.b
}

func (i *squ) sum() int {
	return i.a * 2
}
