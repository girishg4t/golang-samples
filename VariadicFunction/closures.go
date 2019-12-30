package main

import "fmt"

func mainfun() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	callme := mainfun()

	fmt.Println(callme())
	fmt.Println(callme())
	fmt.Println(callme())

	callmeAgain := mainfun()
	fmt.Println(callmeAgain())
}
