package main

import "fmt"

func main() {
	fmt.Println(SumOfNum([]int{1, 2, 3, 4}))
}

func SumOfNum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum

}
