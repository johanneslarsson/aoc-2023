package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := 13
	got := getPartOne(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 21213
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 30
	got := getPartTwo(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 8549735 // too low too low
	got := getPartTwo(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}
