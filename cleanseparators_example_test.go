package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleCleanSeparators() {

	var paths []string = []string{
		"/",
		"//",
		"///",
		"////",

		"something",
		"something/",
		"something//",
		"something///",
		"something////",

		"/////something",
		"////something/",
		"///something//",
		"//something///",
		"/something////",



		"//////once/////twice////thrice///fource//",
	}


	for _, p := range paths {

		cleanedPath := path.CleanSeparators(p)
		fmt.Printf("%q -> %q\n", p, cleanedPath)
	}

	// Output:
	//
	// "/" -> "/"
	// "//" -> "/"
	// "///" -> "/"
	// "////" -> "/"
	// "something" -> "something"
	// "something/" -> "something/"
	// "something//" -> "something/"
	// "something///" -> "something/"
	// "something////" -> "something/"
	// "/////something" -> "/something"
	// "////something/" -> "/something/"
	// "///something//" -> "/something/"
	// "//something///" -> "/something/"
	// "/something////" -> "/something/"
	// "//////once/////twice////thrice///fource//" -> "/once/twice/thrice/fource/"
}
