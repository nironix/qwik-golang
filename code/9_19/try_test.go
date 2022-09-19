package main

import (
	"testing"
)

func TestReval(t *testing.T) {
	str := Reval()
	stre := "hey"
	if str != stre {
		t.Error("nope " + str)
	}
}
