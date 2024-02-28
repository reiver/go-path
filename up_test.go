package path

import (
	"testing"
)

func TestUp(t *testing.T) {

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
			Expected: "/once/",
		},
		{
			Path: "/once/twice/thrice",
			Expected: "/once/twice/",
		},
		{
			Path: "/once/twice/thrice/fource",
			Expected: "/once/twice/thrice/",
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
			Expected: "once/",
		},
		{
			Path: "once/twice/thrice",
			Expected: "once/twice/",
		},
		{
			Path: "once/twice/thrice/fource",
			Expected: "once/twice/thrice/",
		},



		{
			Path: "dragonball.jpeg",
			Expected: "",
		},
		{
			Path: "dragonball.jpeg/picollo",
			Expected: "dragonball.jpeg/",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge",
			Expected: "dragonball.jpeg/picollo/",
		},



		{
			Path: "/one/😈/three.txt",
			Expected: "/one/😈/",
		},
		{
			Path: "one/😈/three.txt",
			Expected: "one/😈/",
		},



		{
			Path: "/path/to/archive.tar.gz",
			Expected: "/path/to/",
		},
		{
			Path: "path/to/archive.tar.gz",
			Expected: "path/to/",
		},



		{
			Path: "/",
			Expected: "/",
		},
		{
			Path: "/something",
			Expected: "/",
		},
		{
			Path: "/something/",
			Expected: "/",
		},
		{
			Path: "/wow",
			Expected: "/",
		},
		{
			Path: "/wow/",
			Expected: "/",
		},
		{
			Path: "/wow/pow",
			Expected: "/wow/",
		},
		{
			Path: "/wow/pow/",
			Expected: "/wow/",
		},
	}

	for testNumber, test := range  tests {

		var path []byte = []byte(test.Path)

		up(&path)

		actual   := string(path)
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
