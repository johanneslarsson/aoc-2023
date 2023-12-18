package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := 62
	got := getPartOne(getRows("../test1.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 36807
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 952408144115
	got := getPartTwo(getRows("../test1.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 48797603984357
	got := getPartTwo(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}
