// +build !windows

package main

// cleanRoot removes the leading slash when running on windows.
func cleanRoot(s string) string {
	return s
}
