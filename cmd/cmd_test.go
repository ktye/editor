// +build windows

package cmd

import "testing"

func TestPath(t *testing.T) {
	testCases := []struct {
		root, name      string
		base, dir, path string
	}{
		{"", "", ".", ".", "."},
		{"", "/alpha.go", "alpha.go", `\`, `\alpha.go`},
		{"d:/abc", "", ".", `d:\abc`, `d:\abc`},
		{"d:/abc/def", "/ghi/alpha.go", `alpha.go`, `d:\abc\def\ghi`, `d:\abc\def\ghi\alpha.go`},
		{"d:/abc/def", "/ghi/alpha.go:23:4", `alpha.go`, `d:\abc\def\ghi`, `d:\abc\def\ghi\alpha.go`},
	}

	for _, tc := range testCases {
		c := Cmd{
			Root: tc.root,
			Name: tc.name,
		}
		if s := c.Base(); s != tc.base {
			t.Fatalf("base expected %s, got %s\n", tc.base, s)
		}
		if s := c.Directory(); s != tc.dir {
			t.Fatalf("dir expected %s, got %s\n", tc.dir, s)
		}
		if s, _ := c.Path(); s != tc.path {
			t.Fatalf("path expected %s, got %s\n", tc.path, s)
		}
	}
}

func TestArgPath(t *testing.T) {
	testCases := []struct {
		root, name, arg string
		path            string
	}{
		{"d:/abc/def", "/ghi/alpha.go:23", `alpha.go`, `d:\abc\def\ghi\alpha.go`},
		{"d:/abc/def", "/ghi/alpha.go:23", `.\alpha.go`, `d:\abc\def\ghi\alpha.go`},
		{"d:/abc/def", "/ghi/alpha.go:23", `..\alpha.go`, `d:\abc\def\alpha.go`},
		{"d:/abc", "/def/", "ghi/", `d:\abc\def\ghi`},
	}

	for _, tc := range testCases {
		c := Cmd{
			Root: tc.root,
			Name: tc.name,
		}
		if s, _ := c.ArgPath(tc.arg); s != tc.path {
			t.Fatalf("arg path expected %s, got %s\n", tc.path, s)
		}
	}
}

func TestTargetPath(t *testing.T) {
	testCases := []struct {
		root, path string
		rel        string
	}{
		{"d:/abc/def", `d:\abc\def\ghi\alpha.go`, `/ghi/alpha.go`},
		{"d:/abc/def", `d:\abc\def\alpha.go`, `/alpha.go`},
		{"d:/abc/def", `c:\Temp\file.xyz`, "c:/Temp/file.xyz"},
		{"d:/abc/def", `d:\abc\def`, "/"},
	}

	for _, tc := range testCases {
		c := Cmd{
			Root: tc.root,
		}
		if s := c.TargetPath(tc.path); s != tc.rel {
			t.Fatalf("target path expected %s, got %s\n", tc.rel, s)
		}
	}
}
