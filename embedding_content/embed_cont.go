package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed text.txt
var text string

func main() {
	strs := strings.Split(text, "\n")
	for _, v := range strs {
		fmt.Println(v)
	}

}
