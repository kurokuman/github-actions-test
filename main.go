package main

import "fmt"

func main() {
	fmt.Println(sumOfNum([]int{1, 3, 4}))
}

func sumOfNum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum

}
