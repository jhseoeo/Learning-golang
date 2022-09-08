package bar

import (
	"fmt"

	"github.com/learning-go/9/internal_package/foo/internal"
)

func Bar() {
	i := internal.Double(2)
	fmt.Println(i, "in Bar")
}
