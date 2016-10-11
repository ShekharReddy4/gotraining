package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_Exec(t *testing.T) {
	bb := &bytes.Buffer{}
	Exec(bb)

	exp := "Hello, World!"
	act := strings.TrimSpace(bb.String())

	if act != exp {
		t.Fatalf("expected %s got %s", exp, act)
	}
}