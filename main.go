package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println(SumOfNum([]int{1, 2, 3, 4}))
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Println(u1)
}

func SumOfNum(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}
	return sum

}
