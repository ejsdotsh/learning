package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	want := 4
	got := count(b, "word")
	if got != want {
		t.Errorf("expected: %d, got: %d\n", want, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	want := 3
	got := count(b, "line")

	if got != want {
		t.Errorf("expected: %d lines, got: %d\n", want, got)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3")
	// bc := scanByteCounter{}

	want := 17
	got := count(b, "bytes")
	if got != want {
		t.Errorf("expected: %d, got: %d\n", want, got)
	}
}
