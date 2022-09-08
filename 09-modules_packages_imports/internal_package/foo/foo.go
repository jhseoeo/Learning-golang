package foo

import (
	"fmt"

	"github.com/learning-go/9/internal_package/foo/internal"
	"github.com/learning-go/9/internal_package/foo/sibling"
)

func Foo() {
	sibling.Sibling()

	i := internal.Double(2)
	fmt.Println(i, "in Foo")
}
