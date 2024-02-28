package path_test

import (
	"testing"

	"github.com/reiver/go-path"
)

func TestSplit(t *testing.T) {

	tests := []struct{
		Path string
		ExpectedDir string
		ExpectedFile string
	}{
		{
			Path: "",
			ExpectedDir: "",
			ExpectedFile: "",
		},



		{
			Path: "/",
			ExpectedDir: "/",
			ExpectedFile: "",
		},



		{
			Path: "photo.jpeg",
			ExpectedDir: "",
			ExpectedFile: "photo.jpeg",
		},
		{
			Path: "/photo.jpeg",
			ExpectedDir: "/",
			ExpectedFile: "photo.jpeg",
		},



		{
			Path: "path/to/photo.jpeg",
			ExpectedDir: "path/to/",
			ExpectedFile: "photo.jpeg",
		},
		{
			Path: "/path/to/photo.jpeg",
			ExpectedDir: "/path/to/",
			ExpectedFile: "photo.jpeg",
		},



		{
			Path: "/../../////.././.././something/../wow/here/it.png",
			ExpectedDir: "/../../////.././.././something/../wow/here/",
			ExpectedFile: "it.png",
		},



		{
			Path: "/apple/banana/cherry/",
			ExpectedDir: "/apple/banana/cherry/",
			ExpectedFile: "",
		},
		{
			Path: "/apple/banana/cherry",
			ExpectedDir: "/apple/banana/",
			ExpectedFile: "cherry",
		},
		{
			Path: "apple/banana/cherry/",
			ExpectedDir: "apple/banana/cherry/",
			ExpectedFile: "",
		},
		{
			Path: "apple/banana/cherry",
			ExpectedDir: "apple/banana/",
			ExpectedFile: "cherry",
		},
	}

	for testNumber, test := range tests {

		actualDir, actualFile := path.Split(test.Path)

		{
			actual   := actualDir
			expected := test.ExpectedDir

			if expected != actual {
				t.Errorf("For test #%d, the actual 'dir' is not what was expected", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUL:    %q", actual)
				t.Logf("PATH: %q", test.Path)
				continue
			}
		}

		{
			actual   := actualFile
			expected := test.ExpectedFile

			if expected != actual {
				t.Errorf("For test #%d, the actual 'file' is not what was expected", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUL:    %q", actual)
				t.Logf("PATH: %q", test.Path)
				continue
			}
		}
	}
}
