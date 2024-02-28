package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleExt() {

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

		extension := path.Ext(p)
		fmt.Printf("the 'extension' of the path %q is %q\n", p, extension)
	}

	// Output:
	//
	// the 'extension' of the path "" is ""
	// the 'extension' of the path "/" is ""
	// the 'extension' of the path "/doc.txt" is ".txt"
	// the 'extension' of the path "doc.txt" is ".txt"
	// the 'extension' of the path "/image.jpeg" is ".jpeg"
	// the 'extension' of the path "image.jpeg" is ".jpeg"
	// the 'extension' of the path "/apple/banana/cherry.html" is ".html"
	// the 'extension' of the path "apple/banana/cherry.html" is ".html"
	// the 'extension' of the path "/path/to/archive.tar.gz" is ".gz"
	// the 'extension' of the path "path/to/archive.tar.gz" is ".gz"
	// the 'extension' of the path "/once/twice/thrice/fource" is ""
	// the 'extension' of the path "once/twice/thrice/fource" is ""
	// the 'extension' of the path "/LICENSE" is ""
	// the 'extension' of the path "LICENSE" is ""
	// the 'extension' of the path "/src/README.md" is ".md"
	// the 'extension' of the path "src/README.md" is ".md"
	// the 'extension' of the path "/ALLCAPS.HTML" is ".HTML"
	// the 'extension' of the path "ALLCAPS.HTML" is ".HTML"
	// the 'extension' of the path "/something/name.dir/filename" is ""
	// the 'extension' of the path "something/name.dir/filename" is ""
}
