package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintNumbers()
}

func PrintNumbers() {
	nums := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			nums <- i
		}
	}()

	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			nums <- j*100 + 1
		}
	}()
	go func() {
		wg.Wait()
		close(nums)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range nums {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}
