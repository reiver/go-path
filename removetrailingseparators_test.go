package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestRemoveTrailingSeparators(t *testing.T) {

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
			Expected: ".",
		},
		{
			Path: ".//",
			Expected: ".",
		},
		{
			Path: ".///",
			Expected: ".",
		},
		{
			Path: ".////",
			Expected: ".",
		},
		{
			Path: "./////",
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
			Path: "..//",
			Expected: "..",
		},
		{
			Path: "..///",
			Expected: "..",
		},
		{
			Path: "..////",
			Expected: "..",
		},
		{
			Path: "../////",
			Expected: "..",
		},



		{
			Path: "something",
			Expected: "something",
		},
		{
			Path: "something/",
			Expected: "something",
		},
		{
			Path: "something//",
			Expected: "something",
		},
		{
			Path: "something///",
			Expected: "something",
		},
		{
			Path: "something////",
			Expected: "something",
		},
		{
			Path: "something/////",
			Expected: "something",
		},



		{
			Path: "one/two//three///",
			Expected: "one/two//three",
		},
		{
			Path: "/one/two//three///",
			Expected: "/one/two//three",
		},



		{
			Path: "/////apple////banana///cherry//date",
			Expected: "/////apple////banana///cherry//date",
		},
		{
			Path: "/////apple////banana///cherry//date/",
			Expected: "/////apple////banana///cherry//date",
		},
		{
			Path: "/////apple////banana///cherry//date//",
			Expected: "/////apple////banana///cherry//date",
		},
		{
			Path: "/////apple////banana///cherry//date///",
			Expected: "/////apple////banana///cherry//date",
		},
		{
			Path: "/////apple////banana///cherry//date////",
			Expected: "/////apple////banana///cherry//date",
		},



		{
			Path: "/////apple////banana///cherry//date/.",
			Expected: "/////apple////banana///cherry//date/.",
		},
		{
			Path: "/////apple////banana///cherry//date/./",
			Expected: "/////apple////banana///cherry//date/.",
		},
		{
			Path: "/////apple////banana///cherry//date/.//",
			Expected: "/////apple////banana///cherry//date/.",
		},
		{
			Path: "/////apple////banana///cherry//date/.///",
			Expected: "/////apple////banana///cherry//date/.",
		},
		{
			Path: "/////apple////banana///cherry//date/.////",
			Expected: "/////apple////banana///cherry//date/.",
		},



		{
			Path: "/////apple////banana///cherry//date/..",
			Expected: "/////apple////banana///cherry//date/..",
		},
		{
			Path: "/////apple////banana///cherry//date/../",
			Expected: "/////apple////banana///cherry//date/..",
		},
		{
			Path: "/////apple////banana///cherry//date/..//",
			Expected: "/////apple////banana///cherry//date/..",
		},
		{
			Path: "/////apple////banana///cherry//date/..///",
			Expected: "/////apple////banana///cherry//date/..",
		},
		{
			Path: "/////apple////banana///cherry//date/..////",
			Expected: "/////apple////banana///cherry//date/..",
		},
	}

	for testNumber, test := range tests {

		actual := path.RemoveTrailingSeparators(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'clear-trailing-slash' result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH: %q", test.Path)
			continue
		}

	}
}
