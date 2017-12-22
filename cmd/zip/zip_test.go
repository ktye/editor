package main

import (
	"bytes"
	"testing"

	"github.com/ktye/editor/cmd"
)

func TestSplitZipPath(t *testing.T) {
	testCases := []struct {
		name string
		path string
		rel  string
	}{
		{"/path/to/file.txt", "", ""},
		{"/path/to/file.zip", "/path/to/file.zip", "/"},
		{"/path/to/file.zip/", "/path/to/file.zip", "/"},
		{"/path/to/file.zip/a.txt", "/path/to/file.zip", "/a.txt"},
		{"/path/to/file.zip/a/b/c/", "/path/to/file.zip", "/a/b/c/"},
	}
	for _, tc := range testCases {
		c := program{
			Cmd: cmd.Cmd{
				Name: tc.name,
			},
		}
		a, b := c.splitZipPath()
		if a != tc.path || b != tc.rel {
			t.Fatalf("%+v: got: '%s, '%s'\n", tc, a, b)
		}
	}
}

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name string
		text string
	}{
		{"/zip/test/file.zip", "README.txt\na/\nb/\n"},
		{"/zip/test/file.zip/", "README.txt\na/\nb/\n"},
		{"/zip/test/file.zip/README.txt", "This zip file is used for go test DO NOT MODIFY!\n"},
		{"/zip/test/file.zip/a/", "aa/\n"},
		{"/zip/test/file.zip/a/aa/", "aa.txt\n"},
		{"/zip/test/file.zip/b/", "b.txt\n"},
	}
	for i, tc := range testCases {
		var in, out bytes.Buffer

		req := cmd.NewTestRequest()
		req.Name = tc.name
		req.Encode(&in)

		expected := cmd.NewTestRequest()
		expected.Name = tc.name
		expected.Tags = ""
		expected.Type = ""
		expected.Clean = true
		expected.Text = tc.text

		var p program
		p.SetIO(&in, &out, nil)
		if err := p.Run(); err != nil {
			t.Fatal(err)
		}
		if err := cmd.CompareTestResults(p.Cmd, *expected); err != nil {
			t.Fatalf("test case %d: %s\n", i, err)
		}
	}
}
