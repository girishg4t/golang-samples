package main

import "fmt"

func pointer(val *int) {
	*val = 0
}

func passbyVal(i int) {
	i = 2
}

func main() {
	i := 1

	passbyVal(i)
	fmt.Println(i)

	pointer(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
