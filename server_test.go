package main

import "testing"

func TestMain(t *testing.T) {
	if 1 != 1 {
		t.Error("Something went terribly wrong!")
	}
}
