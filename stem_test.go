package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestStem(t *testing.T) {

	tests := []struct{
		Path string
		Expected string
	}{
		{
			Path: "",
			Expected: "",
		},



		{
			Path: "/",
			Expected: "",
		},



		{
			Path: "/apple",
			Expected: "apple",
		},
		{
			Path: "/banana",
			Expected: "banana",
		},
		{
			Path: "/cherry",
			Expected: "cherry",
		},

		{
			Path: "/once",
			Expected: "once",
		},
		{
			Path: "/once/twice",
			Expected: "twice",
		},
		{
			Path: "/once/twice/thrice",
			Expected: "thrice",
		},
		{
			Path: "/once/twice/thrice/fource",
			Expected: "fource",
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
			Expected: "twice",
		},
		{
			Path: "once/twice/thrice",
			Expected: "thrice",
		},
		{
			Path: "once/twice/thrice/fource",
			Expected: "fource",
		},



		{
			Path: "dragonball.jpeg",
			Expected: "dragonball",
		},
		{
			Path: "dragonball.jpeg/picollo",
			Expected: "picollo",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge",
			Expected: "dodge",
		},



		{
			Path: "/one/😈/three.txt",
			Expected: "three",
		},
		{
			Path: "one/😈/three.txt",
			Expected: "three",
		},



		{
			Path: "/path/to/archive.tar.gz",
			Expected: "archive.tar",
		},
		{
			Path: "path/to/archive.tar.gz",
			Expected: "archive.tar",
		},
	}

	for testNumber, test := range  tests {

		actual := path.Stem(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for 'stem' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH:     %q", test.Path)
			continue
		}
	}
}
