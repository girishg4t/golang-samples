package main

import (
	"fmt"
	"math/rand"
	"time"
)

//RemoveIndex to remove the element from array
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func shuffleArray(randNumber []int) []int {
	l := len(randNumber)
	newrandArr := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(randNumber); i++ {

		r := rand.Intn(l - i)
		newrandArr = append(newrandArr, randNumber[r])
		RemoveIndex(randNumber, r)
	}

	return newrandArr
}

func main() {

	randNumber := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	fmt.Println(shuffleArray(randNumber))
}
