package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}

	time.Sleep(1 * time.Second)

	manyGoWait()

}

func manyGoWait() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

func hello(i int) {
	println("hello goroutine : " + fmt.Sprint(i))
}
