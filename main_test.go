package main

import (
	"fmt"
	"testing"
)

func TestSumOfNum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	act := SumOfNum(input)
	ext := 15

	if act != ext {
		t.Fatal("error")
	} else {
		fmt.Println(act)
	}
}
