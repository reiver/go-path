package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestExt(t *testing.T) {

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
			Expected: "",
		},
		{
			Path: "/banana",
			Expected: "",
		},
		{
			Path: "/cherry",
			Expected: "",
		},

		{
			Path: "/once",
			Expected: "",
		},
		{
			Path: "/once/twice",
			Expected: "",
		},
		{
			Path: "/once/twice/thrice",
			Expected: "",
		},
		{
			Path: "/once/twice/thrice/fource",
			Expected: "",
		},



		{
			Path: "apple",
			Expected: "",
		},
		{
			Path: "banana",
			Expected: "",
		},
		{
			Path: "cherry",
			Expected: "",
		},

		{
			Path: "once",
			Expected: "",
		},
		{
			Path: "once/twice",
			Expected: "",
		},
		{
			Path: "once/twice/thrice",
			Expected: "",
		},
		{
			Path: "once/twice/thrice/fource",
			Expected: "",
		},



		{
			Path: "dragonball.jpeg",
			Expected: ".jpeg",
		},
		{
			Path: "dragonball.jpeg/picollo",
			Expected: "",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge",
			Expected: "",
		},



		{
			Path: "/one/😈/three.txt",
			Expected: ".txt",
		},
		{
			Path: "one/😈/three.txt",
			Expected: ".txt",
		},



		{
			Path: "/path/to/archive.tar.gz",
			Expected: ".gz",
		},
		{
			Path: "path/to/archive.tar.gz",
			Expected: ".gz",
		},
	}

	for testNumber, test := range  tests {

		actual := path.Ext(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for 'ext' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH:     %q", test.Path)
			continue
		}
	}
}
