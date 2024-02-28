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
			Path: ".",
			Expected: ".",
		},
		{
			Path: "./",
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
