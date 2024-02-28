package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleRemoteTrailingSeparators() {

	var paths []string = []string{
		"/apple/banana/cherry",
		"/apple/banana/cherry/",
		"/apple/banana/cherry//",
		"/apple/banana/cherry///",
	}

	for _, p := range paths {

		var result string = path.RemoveTrailingSeparators(p)

		fmt.Printf("%q -> %q\n", p, result)

	}

	// Output:
	//
	// "/apple/banana/cherry" -> "/apple/banana/cherry"
	// "/apple/banana/cherry/" -> "/apple/banana/cherry"
	// "/apple/banana/cherry//" -> "/apple/banana/cherry"
	// "/apple/banana/cherry///" -> "/apple/banana/cherry"
}
