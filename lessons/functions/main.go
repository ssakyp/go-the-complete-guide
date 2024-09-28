package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, numbers...)
	sum2 := sumup(2, 5, 15)
	fmt.Println(sum, sum2)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, v := range numbers {
		sum += v
	}
	return sum
}
