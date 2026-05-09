package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	testCountWords(t, "word1 word2 word3 word4\n", 4)
	testCountWords(t, "word1 word2 word3\n", 3)
	testCountWords(t, "word1 word2\n", 2)
	testCountWords(t, "word1\n", 1)
	testCountWords(t, "\n", 0)
}

func TestCountLines(t *testing.T) {
	testCountLines(t, "", 0)
	testCountLines(t, "\n", 1)
	testCountLines(t, "line1", 1)
	testCountLines(t, "line1\n", 1)
	testCountLines(t, "\n\n", 2)
	testCountLines(t, "line1\nline2\n", 2)
	testCountLines(t, "line1\nline2\nline3", 3)
}

func TestCountBytes(t *testing.T) {
	testCountBytes(t, "", 0)
	testCountBytes(t, "1", 1)
	testCountBytes(t, "1\n", 2)
}

func testCountWords(t *testing.T, input string, expected int) {
	b := bytes.NewBufferString(input)
	res := count(b, false, false)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func testCountLines(t *testing.T, input string, expected int) {
	b := bytes.NewBufferString(input)
	res := count(b, true, false)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func testCountBytes(t *testing.T, input string, expected int) {
	b := bytes.NewBufferString(input)
	res := count(b, false, true)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
