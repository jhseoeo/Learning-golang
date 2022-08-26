package sibling

import (
	"fmt"

	"github.com/learning-go/9/internal_package/foo/internal"
)

func Sibling() {
	i := internal.Double(2)
	fmt.Println(i, "in Sibling")
}
