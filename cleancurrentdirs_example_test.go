package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleCleanCurrentDirs() {

	var paths []string = []string{
		"/",

		"/.",

		"./",

		"/./",
		"/././",
		"/./././",

		"././",
		"./././",
		"././././",


		"/./.",
		"/././.",
		"/./././.",

		"/././././././././dir/././././././file.txt",
		"././././././././dir/././././././file.txt",
	}


	for _, p := range paths {

		cleanedPath := path.CleanCurrentDirs(p)
		fmt.Printf("%q -> %q\n", p, cleanedPath)
	}

	// Output:
	//
	// "/" -> "/"
	// "/." -> "/"
	// "./" -> "."
	// "/./" -> "/"
	// "/././" -> "/"
	// "/./././" -> "/"
	// "././" -> "."
	// "./././" -> "."
	// "././././" -> "."
	// "/./." -> "/"
	// "/././." -> "/"
	// "/./././." -> "/"
	// "/././././././././dir/././././././file.txt" -> "/dir/file.txt"
	// "././././././././dir/././././././file.txt" -> "dir/file.txt"
}
