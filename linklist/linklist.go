package main

import "fmt"

type node struct {
	val   int
	pNext *node
}

func main() {
	nList := createLinklist()
	printList(nList)
	ll := reverse(nList)
	printList(ll)
}

func createLinklist() *node {
	pHead := &node{1, nil}
	for i := 2; i <= 5; i++ {
		pCurr := &node{i, pHead}
		pHead = pCurr
	}
	return pHead
}

func reverse(l *node) *node {
	var pHead *node = nil
	pCurr := l
	for pCurr.pNext != nil {
		pTemp := pCurr.pNext
		pCurr.pNext = pHead
		pHead = pCurr
		pCurr = pTemp
	}
	pCurr.pNext = pHead
	return pCurr
}

func printList(ll *node) {
	for ll.pNext != nil {
		fmt.Println(ll.val)
		ll = ll.pNext
	}
	fmt.Println(ll.val)
}
