package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestCleanCurrentDirs(t *testing.T) {

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
			Path: "/once",
			Expected: "/once",
		},
		{
			Path: "/once/twice",
			Expected: "/once/twice",
		},
		{
			Path: "/once/twice/thrice",
			Expected: "/once/twice/thrice",
		},
		{
			Path: "/once/twice/thrice/fource",
			Expected: "/once/twice/thrice/fource",
		},



		{
			Path: "once",
			Expected: "once",
		},
		{
			Path: "once/twice",
			Expected: "once/twice",
		},
		{
			Path: "once/twice/thrice",
			Expected: "once/twice/thrice",
		},
		{
			Path: "once/twice/thrice/fource",
			Expected: "once/twice/thrice/fource",
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
			Path: "/././",
			Expected: "/",
		},
		{
			Path: "/./././",
			Expected: "/",
		},
		{
			Path: "/././././",
			Expected: "/",
		},
		{
			Path: "/./././././",
			Expected: "/",
		},
		{
			Path: "/././././././",
			Expected: "/",
		},



		{
			Path: "./",
			Expected: ".",
		},
		{
			Path: "././",
			Expected: ".",
		},
		{
			Path: "./././",
			Expected: ".",
		},
		{
			Path: "././././",
			Expected: ".",
		},
		{
			Path: "./././././",
			Expected: ".",
		},
		{
			Path: "././././././",
			Expected: ".",
		},



		{
			Path: "/./file.txt",
			Expected: "/file.txt",
		},
		{
			Path: "/././file.txt",
			Expected: "/file.txt",
		},
		{
			Path: "/./././file.txt",
			Expected: "/file.txt",
		},
		{
			Path: "/././././file.txt",
			Expected: "/file.txt",
		},
		{
			Path: "/./././././file.txt",
			Expected: "/file.txt",
		},
		{
			Path: "/././././././file.txt",
			Expected: "/file.txt",
		},



		{
			Path: "./file.txt",
			Expected: "file.txt",
		},
		{
			Path: "././file.txt",
			Expected: "file.txt",
		},
		{
			Path: "./././file.txt",
			Expected: "file.txt",
		},
		{
			Path: "././././file.txt",
			Expected: "file.txt",
		},
		{
			Path: "./././././file.txt",
			Expected: "file.txt",
		},
		{
			Path: "././././././file.txt",
			Expected: "file.txt",
		},



		{
			Path: "/./././././././dir/././././././file.txt",
			Expected: "/dir/file.txt",
		},
		{
			Path: "./././././././dir/././././././file.txt",
			Expected: "dir/file.txt",
		},



		{
			Path: "/once/twice/thrice/fource/.",
			Expected: "/once/twice/thrice/fource/",
		},
		{
			Path: "/once/twice/thrice/fource/./",
			Expected: "/once/twice/thrice/fource/",
		},
		{
			Path: "/once/twice/thrice/fource/./.",
			Expected: "/once/twice/thrice/fource/",
		},
		{
			Path: "/once/twice/thrice/fource/././",
			Expected: "/once/twice/thrice/fource/",
		},
		{
			Path: "/once/twice/thrice/fource/././.",
			Expected: "/once/twice/thrice/fource/",
		},
		{
			Path: "/once/twice/thrice/fource/./././",
			Expected: "/once/twice/thrice/fource/",
		},



		{
			Path: "once/twice/thrice/fource/.",
			Expected: "once/twice/thrice/fource/",
		},
		{
			Path: "once/twice/thrice/fource/./",
			Expected: "once/twice/thrice/fource/",
		},
		{
			Path: "once/twice/thrice/fource/./.",
			Expected: "once/twice/thrice/fource/",
		},
		{
			Path: "once/twice/thrice/fource/././",
			Expected: "once/twice/thrice/fource/",
		},
		{
			Path: "once/twice/thrice/fource/././.",
			Expected: "once/twice/thrice/fource/",
		},
		{
			Path: "once/twice/thrice/fource/./././",
			Expected: "once/twice/thrice/fource/",
		},



		{
			Path: "/apple/banana/cherry/./date/",
			Expected: "/apple/banana/cherry/date/",
		},
		{
			Path: "/apple/banana/cherry/././date/",
			Expected: "/apple/banana/cherry/date/",
		},
		{
			Path: "/apple/banana/cherry/./././date/",
			Expected: "/apple/banana/cherry/date/",
		},
		{
			Path: "/apple/banana/cherry/././././date/",
			Expected: "/apple/banana/cherry/date/",
		},



		{
			Path: "/apple/banana/cherry/./date",
			Expected: "/apple/banana/cherry/date",
		},
		{
			Path: "/apple/banana/cherry/././date",
			Expected: "/apple/banana/cherry/date",
		},
		{
			Path: "/apple/banana/cherry/./././date",
			Expected: "/apple/banana/cherry/date",
		},
		{
			Path: "/apple/banana/cherry/././././date",
			Expected: "/apple/banana/cherry/date",
		},



		{
			Path: "apple/banana/cherry/./date/",
			Expected: "apple/banana/cherry/date/",
		},
		{
			Path: "apple/banana/cherry/././date/",
			Expected: "apple/banana/cherry/date/",
		},
		{
			Path: "apple/banana/cherry/./././date/",
			Expected: "apple/banana/cherry/date/",
		},
		{
			Path: "apple/banana/cherry/././././date/",
			Expected: "apple/banana/cherry/date/",
		},


		{
			Path: "apple/banana/cherry/./date",
			Expected: "apple/banana/cherry/date",
		},
		{
			Path: "apple/banana/cherry/././date",
			Expected: "apple/banana/cherry/date",
		},
		{
			Path: "apple/banana/cherry/./././date",
			Expected: "apple/banana/cherry/date",
		},
		{
			Path: "apple/banana/cherry/././././date",
			Expected: "apple/banana/cherry/date",
		},
	}

	for testNumber, test := range  tests {

		actual := path.CleanCurrentDirs(test.Path)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for 'clean-current-dirs' result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PATH:     %q", test.Path)
			continue
		}
	}
}
