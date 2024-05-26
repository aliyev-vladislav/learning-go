package main

import (
	"fmt"
)

func main() {
	processData()
}

func processData() {
	num1 := make(chan int)
	num2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			num1 <- i
		}
		close(num1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			num2 <- i
		}
		close(num2)
	}()

	for count := 0; count < 2; {
		select {
		case v, ok := <-num1:
			if !ok {
				num1 = nil
				count++
				continue
			}
			fmt.Println("Goroutine 1: ", v)
		case v, ok := <-num2:
			if !ok {
				num2 = nil
				count++
				continue
			}
			fmt.Println("Goroutine 2: ", v)
		}
	}
}
