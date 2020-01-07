package main

import "fmt"

func main() {
	reverseLinklist()
	mergeLinklist()
}

func reverseLinklist() {
	nList := createLinklist(1)
	printList(nList)
	ll := reverse(nList)
	printList(ll)
}

func mergeLinklist() {
	nListWith1 := createLinklist(10)
	printList(nListWith1)
	fmt.Println("one done")
	nListWith2 := createLinklist(1)
	printList(nListWith2)
	fmt.Println("two done")
	r := merge(nListWith1, nListWith2)
	printList(r)
}
