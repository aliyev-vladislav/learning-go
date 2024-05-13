package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed english_right.txt
var engRight string

//go:embed russian_right.txt
var rusRight string

//go:embed japan_right.txt
var japRight string

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	switch strings.ToLower(os.Args[1]) {
	case "english":
		fmt.Println(engRight)
	case "russian":
		fmt.Println(rusRight)
	case "japan":
		fmt.Println(japRight)
	default:
		fmt.Println("File not found")
		os.Exit(1)
	}

}
