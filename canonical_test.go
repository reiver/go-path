package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestCanonical(t *testing.T) {

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
			Path: "/.",
			Expected: "/",
		},
		{
			Path: "/./",
			Expected: "/",
		},
		{
			Path: "/./.",
			Expected: "/",
		},
		{
			Path: "/././",
			Expected: "/",
		},
		{
			Path: "/././.",
			Expected: "/",
		},
		{
			Path: "/./././",
			Expected: "/",
		},
		{
			Path: "/./././.",
			Expected: "/",
		},
		{
			Path: "/././././",
			Expected: "/",
		},



		{
			Path: "/..",
			Expected: "/",
		},
		{
			Path: "/../",
			Expected: "/",
		},
		{
			Path: "/../..",
			Expected: "/",
		},
		{
			Path: "/../../",
			Expected: "/",
		},
		{
			Path: "/../../..",
			Expected: "/",
		},
		{
			Path: "/../../../",
			Expected: "/",
		},
		{
			Path: "/../../../..",
			Expected: "/",
		},
		{
			Path: "/../../../../",
			Expected: "/",
		},



		{
			Path: "//",
			Expected: "/",
		},
		{
			Path: "///",
			Expected: "/",
		},
		{
			Path: "////",
			Expected: "/",
		},
		{
			Path: "/////",
			Expected: "/",
		},



		{
			Path: ".",
			Expected: ".",
		},
		{
			Path: "./",
			Expected: ".",
		},



		{
			Path: "./.",
			Expected: ".",
		},
		{
			Path: "././",
			Expected: ".",
		},
		{
			Path: "././.",
			Expected: ".",
		},
		{
			Path: "./././",
			Expected: ".",
		},
		{
			Path: "./././.",
			Expected: ".",
		},
		{
			Path: "././././",
			Expected: ".",
		},



		{
			Path: "..",
			Expected: "..",
		},
		{
			Path: "../",
			Expected: "..",
		},



		{
			Path: "../..",
			Expected: "../..",
		},
		{
			Path: "../../",
			Expected: "../..",
		},



		{
			Path: "../../..",
			Expected: "../../..",
		},
		{
			Path: "../../../",
			Expected: "../../..",
		},



		{
			Path: "../../../..",
			Expected: "../../../..",
		},
		{
			Path: "../../../../",
			Expected: "../../../..",
		},



		{
			Path: "apple////banana///cherry//",
			Expected: "apple/banana/cherry/",
		},
		{
			Path: "/apple////banana///cherry//",
			Expected: "/apple/banana/cherry/",
		},
		{
			Path: "//apple////banana///cherry//",
			Expected: "/apple/banana/cherry/",
		},
		{
			Path: "///apple////banana///cherry//",
			Expected: "/apple/banana/cherry/",
		},
		{
			Path: "////apple////banana///cherry//",
			Expected: "/apple/banana/cherry/",
		},



		{
			Path: "apple////banana///cherry",
			Expected: "apple/banana/cherry",
		},
		{
			Path: "/apple////banana///cherry",
			Expected: "/apple/banana/cherry",
		},
		{
			Path: "//apple////banana///cherry",
			Expected: "/apple/banana/cherry",
		},
		{
			Path: "///apple////banana///cherry",
			Expected: "/apple/banana/cherry",
		},
		{
			Path: "////apple////banana///cherry",
			Expected: "/apple/banana/cherry",
		},



		{
			Path: "/../../..////./././../something/././/..///./wow/././pow",
			Expected: "/wow/pow",
		},
		{
			Path: "/../../..////./././../something/././/..///./wow/././pow/",
			Expected: "/wow/pow/",
		},
		{
			Path: "/../../..////./././../something/././/..///./wow/././pow/.",
			Expected: "/wow/pow/",
		},
		{
			Path: "/../../..////./././../something/././/..///./wow/././pow/./",
			Expected: "/wow/pow/",
		},
		{
			Path: "/../../..////./././../something/././/..///./wow/././pow/./.",
			Expected: "/wow/pow/",
		},



		{
			Path: "////../../..//../../home//username/",
			Expected: "/home/username/",
		},
		{
			Path: "////../../..//../../home//username",
			Expected: "/home/username",
		},



		{
			Path: "/apple",
			Expected: "/apple",
		},
		{
			Path: "/banana",
			Expected: "/banana",
		},
		{
			Path: "/cherry",
			Expected: "/cherry",
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
			Path: "/once",
			Expected: "/once",
		},
		{
			Path: "/once/",
			Expected: "/once/",
		},
		{
			Path: "/once/twice",
			Expected: "/once/twice",
		},
		{
			Path: "/once/twice/",
			Expected: "/once/twice/",
		},
		{
			Path: "/once/twice/thrice",
			Expected: "/once/twice/thrice",
		},
		{
			Path: "/once/twice/thrice/",
			Expected: "/once/twice/thrice/",
		},
		{
			Path: "/once/twice/thrice/fource",
			Expected: "/once/twice/thrice/fource",
		},
		{
			Path: "/once/twice/thrice/fource/",
			Expected: "/once/twice/thrice/fource/",
		},




		{
			Path: "once",
			Expected: "once",
		},
		{
			Path: "once/",
			Expected: "once/",
		},
		{
			Path: "once/twice",
			Expected: "once/twice",
		},
		{
			Path: "once/twice/",
			Expected: "once/twice/",
		},
		{
			Path: "once/twice/thrice",
			Expected: "once/twice/thrice",
		},
		{
			Path: "once/twice/thrice/",
			Expected: "once/twice/thrice/",
		},
		{
			Path: "once/twice/thrice/fource",
			Expected: "once/twice/thrice/fource",
		},
		{
			Path: "once/twice/thrice/fource/",
			Expected: "once/twice/thrice/fource/",
		},



		{
			Path: "/dragonball.jpeg/",
			Expected: "/dragonball.jpeg/",
		},
		{
			Path: "/dragonball.jpeg",
			Expected: "/dragonball.jpeg",
		},
		{
			Path: "/dragonball.jpeg/picollo/",
			Expected: "/dragonball.jpeg/picollo/",
		},
		{
			Path: "/dragonball.jpeg/picollo",
			Expected: "/dragonball.jpeg/picollo",
		},
		{
			Path: "/dragonball.jpeg/picollo/dodge/",
			Expected: "/dragonball.jpeg/picollo/dodge/",
		},
		{
			Path: "/dragonball.jpeg/picollo/dodge",
			Expected: "/dragonball.jpeg/picollo/dodge",
		},

		{
			Path: "dragonball.jpeg/",
			Expected: "dragonball.jpeg/",
		},
		{
			Path: "dragonball.jpeg",
			Expected: "dragonball.jpeg",
		},
		{
			Path: "dragonball.jpeg/picollo/",
			Expected: "dragonball.jpeg/picollo/",
		},
		{
			Path: "dragonball.jpeg/picollo",
			Expected: "dragonball.jpeg/picollo",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge/",
			Expected: "dragonball.jpeg/picollo/dodge/",
		},
		{
			Path: "dragonball.jpeg/picollo/dodge",
			Expected: "dragonball.jpeg/picollo/dodge",
		},



		{
			Path: "/one/😈/three.txt",
			Expected: "/one/😈/three.txt",
		},
		{
			Path: "one/😈/three.txt",
			Expected: "one/😈/three.txt",
		},



		{
			Path: "/path/to/archive.tar.gz",
			Expected: "/path/to/archive.tar.gz",
		},
		{
			Path: "path/to/archive.tar.gz",
			Expected: "path/to/archive.tar.gz",
		},
	}

	for testNumber, test := range  tests {

		actual := path.Canonical(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for 'canonical' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH:     %q", test.Path)
			continue
		}
	}
}
