package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := 6
	got := getPartOne(getRows("../test1.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 20659
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 6
	got := getPartTwo(getRows("../test2.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 15690466351717 //10962774237010832183
	//107541 // wrong
	//134142 too low
	got := getPartTwo(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}
