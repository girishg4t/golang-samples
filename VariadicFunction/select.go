package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan string)
	b := make(chan string)
	start := time.Now()
	go func() {
		time.Sleep(2 * time.Second)
		a <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		b <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-a:
			fmt.Println("received", msg1)
		case msg2 := <-b:
			fmt.Println("received", msg2)
		}
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
