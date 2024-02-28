package path_test

import (
	"fmt"

	"github.com/reiver/go-path"
)

func ExampleJoin() {

	var list [][]string = [][]string{
		[]string{},
		[]string{""},
		[]string{"", ""},

		[]string{"/"},
		[]string{"/", "one"},
		[]string{"/", "one", "two"},
		[]string{"/", "one", "two", "three"},

		[]string{"apple"},
		[]string{"apple", "banana"},
		[]string{"apple", "banana", "cherry"},

		[]string{"/path/to", "file.txt"},
		[]string{"/path/to/", "file.txt"},

		[]string{"/once/twice/thrice/fource.png"},
		[]string{"/once/twice/thrice/", "fource.png"},
		[]string{"/once/twice/thrice", "fource.png"},
		[]string{"/once/twice/", "thrice/fource.png"},
		[]string{"/once/twice", "thrice/fource.png"},
	}

	for _, parts := range list {

		joined := path.Join(parts...)

		printResult(joined, parts...)
	}

	// Output:
	//
	// path.Join() -> "."
	// path.Join("") -> "."
	// path.Join("", "") -> "."
	// path.Join("/") -> "/"
	// path.Join("/", "one") -> "/one"
	// path.Join("/", "one", "two") -> "/one/two"
	// path.Join("/", "one", "two", "three") -> "/one/two/three"
	// path.Join("apple") -> "apple"
	// path.Join("apple", "banana") -> "apple/banana"
	// path.Join("apple", "banana", "cherry") -> "apple/banana/cherry"
	// path.Join("/path/to", "file.txt") -> "/path/to/file.txt"
	// path.Join("/path/to/", "file.txt") -> "/path/to/file.txt"
	// path.Join("/once/twice/thrice/fource.png") -> "/once/twice/thrice/fource.png"
	// path.Join("/once/twice/thrice/", "fource.png") -> "/once/twice/thrice/fource.png"
	// path.Join("/once/twice/thrice", "fource.png") -> "/once/twice/thrice/fource.png"
	// path.Join("/once/twice/", "thrice/fource.png") -> "/once/twice/thrice/fource.png"
	// path.Join("/once/twice", "thrice/fource.png") -> "/once/twice/thrice/fource.png"
}

func printResult(joined string, parts ...string) {

	fmt.Print("path.Join(")
	for i, part := range parts {
		if 0 < i {
			fmt.Print(", ")
		}
		fmt.Printf("%q", part)
	}
	fmt.Printf(") -> %q\n", joined)

}
