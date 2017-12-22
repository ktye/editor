package main

import (
	"bytes"
	"testing"

	"github.com/ktye/editor/cmd"
)

func TestRead(t *testing.T) {

	testCases := []struct {
		name, tags, typ, text string
	}{
		{"/read/test/", "", "text", "../\na/\nfile1\nfile2\n"},
		{"/read/test/a/", "", "text", "../\nfile.go\n"},
		{"/read/test/a/file.go", "-def -doc -Fmt -Install", "text/go", "package a\n"},
	}

	for i, tc := range testCases {
		var in, out bytes.Buffer

		req := cmd.NewTestRequest()
		req.Name = tc.name
		req.Encode(&in)

		expected := cmd.NewTestRequest()
		expected.Name = tc.name
		expected.Tags = tc.tags
		expected.Type = tc.typ
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

func TestReadArg(t *testing.T) {

	testCases := []struct {
		inName, arg, outName, tags, typ, text string
	}{
		{"/read/test/", "a/file.go", "/read/test/a/file.go", "-def -doc -Fmt -Install", "text/go", "package a\n"},
		{"/read/test/", `.\a\file.go`, "/read/test/a/file.go", "-def -doc -Fmt -Install", "text/go", "package a\n"},
		{"/read/test/", `..\test\a\file.go`, "/read/test/a/file.go", "-def -doc -Fmt -Install", "text/go", "package a\n"},
	}

	for i, tc := range testCases {
		var in, out bytes.Buffer

		req := cmd.NewTestRequest()
		req.Name = tc.inName
		req.Encode(&in)

		expected := cmd.NewTestRequest()
		expected.Name = tc.outName
		expected.Tags = tc.tags
		expected.Type = tc.typ
		expected.Clean = true
		expected.Text = tc.text

		var p program
		p.SetIO(&in, &out, []string{tc.arg})
		if err := p.Run(); err != nil {
			t.Fatal(err)
		}
		if err := cmd.CompareTestResults(p.Cmd, *expected); err != nil {
			t.Fatalf("test case %d: %s\n", i, err)
		}
	}
}
