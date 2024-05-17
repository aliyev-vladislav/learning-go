package main

import (
	"fmt"
	"sync"
)

func main() {
	result := Parse("hello")
	fmt.Println(result)
	result2 := Parse("goodbye")
	fmt.Println(result2)
}

type SlowComplicatedParser interface {
	Parse(string) string
}

var initParserCached func() SlowComplicatedParser = sync.OnceValue(initParser)

func Parse(dataToParse string) string {
	parser := initParserCached()
	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	fmt.Println("initializing!")
	return SCPI{}
}

type SCPI struct {
}

func (s SCPI) Parse(in string) string {
	if len(in) > 1 {
		return in[0:1]
	}
	return ""
}
