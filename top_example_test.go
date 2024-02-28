package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleTop() {

	var paths []string = []string{
		"",
		"/",
		"/doc.txt",
		"doc.txt",
		"/image.jpeg",
		"image.jpeg",
		"/apple/banana/cherry.html",
		"apple/banana/cherry.html",
		"/path/to/archive.tar.gz",
		"path/to/archive.tar.gz",
		"/once/twice/thrice/fource",
		"once/twice/thrice/fource",
		"/LICENSE",
		"LICENSE",
		"/src/README.md",
		"src/README.md",
		"/ALLCAPS.HTML",
		"ALLCAPS.HTML",
		"/something/name.dir/filename",
		"something/name.dir/filename",
	}


	for _, p := range paths {

		extension := path.Top(p)
		fmt.Printf("the 'top' of the path %q is %q\n", p, extension)
	}

	// Output:
	//
	// the 'top' of the path "" is ""
	// the 'top' of the path "/" is "/"
	// the 'top' of the path "/doc.txt" is "/"
	// the 'top' of the path "doc.txt" is "doc.txt"
	// the 'top' of the path "/image.jpeg" is "/"
	// the 'top' of the path "image.jpeg" is "image.jpeg"
	// the 'top' of the path "/apple/banana/cherry.html" is "/"
	// the 'top' of the path "apple/banana/cherry.html" is "apple"
	// the 'top' of the path "/path/to/archive.tar.gz" is "/"
	// the 'top' of the path "path/to/archive.tar.gz" is "path"
	// the 'top' of the path "/once/twice/thrice/fource" is "/"
	// the 'top' of the path "once/twice/thrice/fource" is "once"
	// the 'top' of the path "/LICENSE" is "/"
	// the 'top' of the path "LICENSE" is "LICENSE"
	// the 'top' of the path "/src/README.md" is "/"
	// the 'top' of the path "src/README.md" is "src"
	// the 'top' of the path "/ALLCAPS.HTML" is "/"
	// the 'top' of the path "ALLCAPS.HTML" is "ALLCAPS.HTML"
	// the 'top' of the path "/something/name.dir/filename" is "/"
	// the 'top' of the path "something/name.dir/filename" is "something"
}
