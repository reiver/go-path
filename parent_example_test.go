package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleParent() {

	var paths []string = []string{
		"/",
		".",
		"..",
		"../..",
		"/image.jpeg",
		"image.jpeg",
		"/apple/banana/cherry.html",
		"apple/banana/cherry.html",
	}


	for _, p := range paths {

		extension := path.Parent(p)
		fmt.Printf("%q -> %q\n", p, extension)
	}

	// Output:
	//
	// "/" -> "/"
	// "." -> ".."
	// ".." -> "../.."
	// "../.." -> "../../.."
	// "/image.jpeg" -> "/"
	// "image.jpeg" -> "."
	// "/apple/banana/cherry.html" -> "/apple/banana/"
	// "apple/banana/cherry.html" -> "apple/banana/"
}
