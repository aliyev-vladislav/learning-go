package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	sum := 0
	count := 0
	for {
		select {
		case <-timeout.Done():
			fmt.Println("sum:", sum, "number of iterations:", count, timeout.Err())
			return
		default:
		}
		num := rand.Intn(100_000_000)
		if num == 1_234 {
			fmt.Println("sum:", sum, "number of iteration:", count, "got 1234")
			return
		}
		sum += num
		count++
	}

}
