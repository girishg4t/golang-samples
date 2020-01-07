package main

import "fmt"

//Node represent item in linklist
type node struct {
	val   int
	pNext *node
}

func createLinklist(mult int) *node {
	pHead := &node{mult, nil}
	for i := 2; i <= 5; i++ {
		pCurr := &node{i * mult, pHead}
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

//For example, if first list is 5->7->17->13->11 and second is 12->10->2->4->6,
//the first list should become 5->12->7->10->17->2->13->4->11->6 and second
//list should become empty.

func merge(flist *node, slist *node) *node {
	pHead := flist

	for flist != nil {
		fTemp := flist.pNext
		flist.pNext = slist
		sTemp := slist.pNext
		slist.pNext = fTemp
		flist = fTemp
		slist = sTemp
	}
	return pHead
}
