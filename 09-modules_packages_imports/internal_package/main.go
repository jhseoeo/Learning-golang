package main

import (
	"fmt"

	"github.com/learning-go/9/internal_package/bar"
	"github.com/learning-go/9/internal_package/foo"
	"github.com/learning-go/9/internal_package/foo/internal"
)

func main() {
	foo.Foo()
	bar.Bar()

	i := internal.Double(2)
	fmt.Println(i, "in main")
}
