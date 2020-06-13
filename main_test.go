package main

import "testing"

func TestsumOfNum(t *testing.T) {
	input := []int{1, 2, 3}
	act := sumOfNum(input)
	ext := 6

	if act != ext {
		t.Fatal("error")
	}
}
