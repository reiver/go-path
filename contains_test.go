package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestContains(t *testing.T) {

	tests := []struct{
		Path string
		Candidate string
		Expected bool
	}{
		{
			Path:      "/",
			Candidate: "/",
			Expected: true,
		},



		{
			Path:      "/once",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/twice",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/twice/",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/fource",
			Candidate: "/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/fource/",
			Candidate: "/",
			Expected: true,
		},



		{
			Path:      "/once",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/twice",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/twice/",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/twice/thrice",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/twice/thrice/",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/twice/thrice/fource",
			Candidate: "/on",
			Expected: false,
		},
		{
			Path:      "/once/twice/thrice/fource/",
			Candidate: "/on",
			Expected: false,
		},



		{
			Path:      "/once",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/twice",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/twice/",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/fource",
			Candidate: "/once",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/fource/",
			Candidate: "/once",
			Expected: true,
		},



		{
			Path:      "/once",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/twice",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/twice/",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/fource",
			Candidate: "/once/",
			Expected: true,
		},
		{
			Path:      "/once/twice/thrice/fource/",
			Candidate: "/once/",
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := path.Contains(test.Candidate, test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'contains' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("PATH:      %q", test.Path)
			t.Logf("CANDIDATE: %q", test.Candidate)
			continue
		}
	}
}
