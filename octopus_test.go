package main

import "testing"

// TODO : grab a test lib
func TestInitFlags(t *testing.T) {
	pattern, repository := InitFlags()
	if *pattern != "" {
		t.Error("flag -pattern should default to \"\"")
	}
	if *repository != "." {
		t.Error("flag -repository should default to \".\"")
	}
}

