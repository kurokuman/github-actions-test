package main

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestSumOfNum(t *testing.T) {
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Println(u1)

	input := []int{1, 2, 3, 4, 5}
	act := SumOfNum(input)
	ext := 15

	if act != ext {
		t.Fatal("error")
	} else {
		fmt.Println(act)
	}
}
