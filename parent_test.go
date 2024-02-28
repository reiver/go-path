package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestParent(t *testing.T) {

	tests := []struct{
		Path string
		Expected string
	}{
		{
			Path: "",
			Expected: "..",
		},



		{
			Path: ".",
			Expected: "..",
		},



		{
			Path: "..",
			Expected: "../..",
		},
		{
			Path: "../..",
			Expected: "../../..",
		},
		{
			Path: "../../..",
			Expected: "../../../..",
		},
		{
			Path: "../../../..",
			Expected: "../../../../..",
		},
		{
			Path: "../../../../..",
			Expected: "../../../../../..",
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
			Expected: ".",
		},
		{
			Path: "banana",
			Expected: ".",
		},
		{
			Path: "cherry",
			Expected: ".",
		},

		{
			Path: "once",
			Expected: ".",
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
			Expected: ".",
		},
		{
			Path: "dragonball.jpeg/",
			Expected: ".",
		},
		{
			Path: "dragonball.jpeg/picollo",
			Expected: "dragonball.jpeg/",
		},
		{
			Path: "dragonball.jpeg/picollo/",
			Expected: "dragonball.jpeg/",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge",
			Expected: "dragonball.jpeg/picollo/",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge/",
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



		{
			Path: "thing",
			Expected: ".",
		},
		{
			Path: "../thing",
			Expected: "..",
		},
		{
			Path: "../../thing",
			Expected: "../..",
		},
		{
			Path: "../../../thing",
			Expected: "../../..",
		},
		{
			Path: "../../../../thing",
			Expected: "../../../..",
		},
	}

	for testNumber, test := range  tests {

		actual := path.Parent(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for 'parent' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH:     %q", test.Path)
			continue
		}
	}
}
