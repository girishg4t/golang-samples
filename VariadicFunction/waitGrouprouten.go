package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Working %d done\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}
