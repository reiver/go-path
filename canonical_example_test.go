package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleCanonical() {

	var paths []string = []string{
		"",
		"/",
		".",
		"..",
		"/../../..////./././../something/././/..///./wow/././pow/./.",
	}


	for _, p := range paths {

		canonical := path.Canonical(p)
		fmt.Printf("%q -> %q\n", p, canonical)
	}

	// Output:
	//
	// "" -> "."
	// "/" -> "/"
	// "." -> "."
	// ".." -> ".."
	// "/../../..////./././../something/././/..///./wow/././pow/./." -> "/wow/pow/"
}
