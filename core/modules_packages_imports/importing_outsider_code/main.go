package main

import "fmt"

type account struct {
	value int
}

func main() {
	s1 := make([]account, 2, 3)
	s2 := append(s1, account{})
	s3 := append(s2, account{})

	acc := &s2[0]
	acc2 := &s3[0]

	acc.value = 100
	acc2.value = 200

	fmt.Println(s1, s2, s3)

	acc.value = 123

	fmt.Println(s1, s2, s3)

}
