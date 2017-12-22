package main

// cleanRoot removes the leading slash from the root.
// root is passed on the URL as ADDR:PORTROOT and starts with a slash.
// e.g. localhost:2017/d:/path/to/root
func cleanRoot(s string) string {
	if len(s) > 0 && s[0] == '/' {
		return s[1:]
	}
	return s
}
