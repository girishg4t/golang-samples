package main

import (
	"fmt"

	p "github.com/girishg4t/golang-samples/stringer/painkiller"
	"github.com/girishg4t/greet"
)

func main() {
	p1 := p.PillType{Pill: 3}
	fmt.Println(p1.String())
	greet.Hello()
}
