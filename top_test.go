package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestTop(t *testing.T) {

	tests := []struct{
		Path string
		Expected string
	}{
		{
			Path: "",
			Expected: ".",
		},



		{
			Path: "/",
			Expected: "/",
		},



		{
			Path: "/apple",
			Expected: "/",
		},
		{
			Path: "/banana",
			Expected: "/",
		},
		{
			Path: "/cherry",
			Expected: "/",
		},

		{
			Path: "/once",
			Expected: "/",
		},
		{
			Path: "/once/twice",
			Expected: "/",
		},
		{
			Path: "/once/twice/thrice",
			Expected: "/",
		},
		{
			Path: "/once/twice/thrice/fource",
			Expected: "/",
		},



		{
			Path: "apple",
			Expected: "apple",
		},
		{
			Path: "banana",
			Expected: "banana",
		},
		{
			Path: "cherry",
			Expected: "cherry",
		},

		{
			Path: "once",
			Expected: "once",
		},
		{
			Path: "once/twice",
			Expected: "once",
		},
		{
			Path: "once/twice/thrice",
			Expected: "once",
		},
		{
			Path: "once/twice/thrice/fource",
			Expected: "once",
		},



		{
			Path: "dragonball.jpeg",
			Expected: "dragonball.jpeg",
		},
		{
			Path: "dragonball.jpeg/picollo",
			Expected: "dragonball.jpeg",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge",
			Expected: "dragonball.jpeg",
		},



		{
			Path: "/one/😈/three.txt",
			Expected: "/",
		},
		{
			Path: "one/😈/three.txt",
			Expected: "one",
		},



		{
			Path: "/path/to/archive.tar.gz",
			Expected: "/",
		},
		{
			Path: "path/to/archive.tar.gz",
			Expected: "path",
		},



		{
			Path: "/////some/path/to/a/file.html",
			Expected: "/",
		},



		{
			Path: "./some/path/to/a/file.html",
			Expected: "some",
		},
		{
			Path: "././some/path/to/a/file.html",
			Expected: "some",
		},
		{
			Path: "./././some/path/to/a/file.html",
			Expected: "some",
		},
		{
			Path: "././././some/path/to/a/file.html",
			Expected: "some",
		},



		{
			Path: "/something",
			Expected: "/",
		},
		{
			Path: "/../something",
			Expected: "/",
		},
		{
			Path: "/../../something",
			Expected: "/",
		},
		{
			Path: "/../../../something",
			Expected: "/",
		},
		{
			Path: "/../../../../something",
			Expected: "/",
		},

		{
			Path: "/something/",
			Expected: "/",
		},
		{
			Path: "/../something/",
			Expected: "/",
		},
		{
			Path: "/../../something/",
			Expected: "/",
		},
		{
			Path: "/../../../something/",
			Expected: "/",
		},
		{
			Path: "/../../../../something/",
			Expected: "/",
		},
	}

	for testNumber, test := range  tests {

		actual := path.Top(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for 'top' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH:     %q", test.Path)
			continue
		}
	}
}
