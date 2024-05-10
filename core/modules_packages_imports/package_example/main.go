package main

import (
	"fmt"

	format "github.com/aliyev-vladislav/learning-go/core/modules_packages_imports/package_example/do-format"
	"github.com/aliyev-vladislav/learning-go/core/modules_packages_imports/package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
