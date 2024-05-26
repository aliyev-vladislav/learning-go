package main

import (
	"fmt"
	"math"
	"sync"
)

var initMap = sync.OnceValue(ProcessData)

func main() {
	m := initMap()
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(m[i])
	}

}

func ProcessData() map[int]float64 {

	m := make(map[int]float64)
	for i := 0; i < 100_000; i++ {
		m[i] = math.Sqrt(float64(i))
	}
	return m
}
