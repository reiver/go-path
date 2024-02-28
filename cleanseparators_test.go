package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestCleanSeparators(t *testing.T) {

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
			Expected: "./",
		},
		{
			Path: ".//",
			Expected: "./",
		},
		{
			Path: ".///",
			Expected: "./",
		},
		{
			Path: ".////",
			Expected: "./",
		},
		{
			Path: "./////",
			Expected: "./",
		},



		{
			Path: "..",
			Expected: "..",
		},
		{
			Path: "../",
			Expected: "../",
		},
		{
			Path: "..//",
			Expected: "../",
		},
		{
			Path: "..///",
			Expected: "../",
		},
		{
			Path: "..////",
			Expected: "../",
		},
		{
			Path: "../////",
			Expected: "../",
		},



		{
			Path: "something",
			Expected: "something",
		},
		{
			Path: "something/",
			Expected: "something/",
		},
		{
			Path: "something//",
			Expected: "something/",
		},
		{
			Path: "something///",
			Expected: "something/",
		},
		{
			Path: "something////",
			Expected: "something/",
		},
		{
			Path: "something/////",
			Expected: "something/",
		},



		{
			Path: "one/two//three///four////",
			Expected: "one/two/three/four/",
		},
		{
			Path: "/one/two//three///four////",
			Expected: "/one/two/three/four/",
		},
	}

	for testNumber, test := range tests {

		actual := path.CleanSeparators(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'clean-separators' result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH: %q", test.Path)
			continue
		}

	}
}
